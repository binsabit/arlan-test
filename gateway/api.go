package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"time"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/julienschmidt/httprouter"
)

var rchans = make(map[string](chan storage.StoreImageResponse))

type gatewayAPI struct {
	port string
}

func NewGatewayAPI(port string) *gatewayAPI {
	return &gatewayAPI{
		port: port,
	}
}

func (g gatewayAPI) runApi() error {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/image/:id", g.handleGetImage)
	router.HandlerFunc(http.MethodPost, "/convert", g.maxBytes(g.handleConvertImage))
	router.HandlerFunc(http.MethodGet, "/", g.handleMainPage)
	router.ServeFiles("/static/*filepath", http.Dir("public"))
	server := &http.Server{
		Addr:    g.port,
		Handler: router,
	}

	return server.ListenAndServe()

}

func (g gatewayAPI) handleGetImage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from get image"))

}

func (g gatewayAPI) handleConvertImage(w http.ResponseWriter, r *http.Request) {
	log.Println("hello from convert")
	//reading from form and saving image to disk
	r.ParseMultipartForm(5 << 20)
	f, fh, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	defer f.Close()
	filename := fh.Filename

	filepath := path.Join("./temp", r.FormValue("imagename"))
	_ = os.Mkdir(filepath, os.ModePerm)
	file, err := os.OpenFile(path.Join(filepath, filename), os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	io.Copy(file, f)
	defer file.Close()
	//converting image to  webp
	err = convertJpegToWebp(path.Join(filepath, filename))
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	//sending to rabiitmq
	img := &storage.Image{
		Uid:  uid(),
		Name: r.FormValue("imagename"),
		Path: path.Join(filepath, filename),
	}
	storeMsg := &storage.StoreImageRequest{
		Uid:     uid(),
		Image:   img,
		ReplyTo: "gateway",
	}

	log.Printf("INFO: convert request: %v", storeMsg)

	//create channel andd to rchans with id
	rchan := make(chan storage.StoreImageResponse)
	rchans[storeMsg.Uid] = rchan

	msg := RabbitMsg{
		QueueName: "storage",
		Message:   *storeMsg,
	}
	pchan <- msg
	waitReply(storeMsg.Uid, rchan, w)
}

func waitReply(uid string, rchan chan storage.StoreImageResponse, w http.ResponseWriter) {
	for {
		select {
		case storeResponse := <-rchan:
			// responses received
			log.Printf("INFO: received reply: %v uid: %s", storeResponse, uid)

			// send response back to client
			response(w, "Created", 201)

			// remove channel from rchans
			delete(rchans, uid)
			return
		case <-time.After(10 * time.Second):
			// timeout
			log.Printf("ERROR: request timeout uid: %s", uid)

			// send response back to client
			response(w, "Timeout", 408)

			// remove channel from rchans
			delete(rchans, uid)
			return
		}

	}
}

func (g gatewayAPI) handleMainPage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./static/forms.html")
	if err != nil {
		log.Println(err)
	}
	t.Execute(w, nil)
}

func (g gatewayAPI) maxBytes(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.Body = http.MaxBytesReader(w, r.Body, 32768+5000000)
		if err := r.ParseForm(); err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		f(w, r)
	}
}

func uid() string {
	t := time.Now().UnixNano() / int64(time.Millisecond)
	return "ops" + strconv.FormatInt(t, 10)
}

func timestamp() int64 {
	return time.Now().UnixNano() / int64(time.Millisecond)
}

func response(w http.ResponseWriter, resp string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, resp)
}
