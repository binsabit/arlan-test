package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"time"

	storage "github.com/binsabit/arlan-test/proto"
	"github.com/julienschmidt/httprouter"
)

var srchans = make(map[string](chan storage.StoreImageResponse))
var grchans = make(map[string](chan storage.GetImageResponse))

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
	router.HandlerFunc(http.MethodGet, "/image/:name", g.handleGetImage)
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

	params := httprouter.ParamsFromContext(r.Context())
	name := params.ByName("name")
	log.Printf("INFO: get image request to image: %s", name)

	log.Println(name)
	getImgMsg := &storage.GetImageRequest{Name: name, ReplyTo: "gateway"}
	log.Printf("INFO: get image request to storage")

	grchan := make(chan storage.GetImageResponse)
	grchans[getImgMsg.Name] = grchan

	msg := RabbitGetImgMsg{
		QueueName: "image",
		Message:   *getImgMsg,
	}
	gchan <- msg
	waitReplyGet(getImgMsg.Name, grchan, w)

}

func (g gatewayAPI) handleConvertImage(w http.ResponseWriter, r *http.Request) {
	log.Printf("INFO: convert imagerequest")
	//reading from form and saving image to disk
	r.ParseMultipartForm(5 << 20)
	f, _, err := r.FormFile("file")
	if err != nil {
		log.Printf("ERROR: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}
	defer f.Close()
	//turning file bute to implement io.Reader for converter
	var buf bytes.Buffer
	if _, err := io.Copy(&buf, f); err != nil {
		log.Println("ERROR: failed to convert multipart file to bytes", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	//converting image to  webp
	res, err := convertJpegToWebp(&buf)
	if err != nil {
		log.Printf("ERROR: error while converting %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	img := &storage.Image{
		Uid:     uid(),
		Name:    r.FormValue("imagename"),
		Content: res,
	}
	storeMsg := &storage.StoreImageRequest{
		Uid:     uid(),
		Image:   img,
		ReplyTo: "gateway",
	}

	log.Printf("INFO: convert request to storage service, image UID %v", storeMsg.Uid)

	//create channel andd to sotorage recive with id
	srchan := make(chan storage.StoreImageResponse)
	srchans[storeMsg.Uid] = srchan

	msg := RabbitStoreMsg{
		QueueName: "storage",
		Message:   *storeMsg,
	}
	schan <- msg
	waitReplyStore(storeMsg.Uid, srchan, w)
}

func waitReplyGet(name string, grchan chan storage.GetImageResponse, w http.ResponseWriter) {
	for {
		select {
		case getResponse := <-grchan:
			// responses received
			log.Printf("INFO: received reply from getimage for name: %s status %v", getResponse.Status, name)

			w.Header().Add("Content-Type", "image/webp")
			w.Write(getResponse.Image.Content)

			delete(grchans, name)
			return
		case <-time.After(15 * time.Second):
			// timeout
			log.Printf("ERROR: get image request timeout uid: %s", uid)

			// send response back to client
			response(w, "Timeout", http.StatusRequestTimeout)

			// remove channel from srchans
			delete(grchans, name)
			return
		}

	}
}

func waitReplyStore(uid string, srchan chan storage.StoreImageResponse, w http.ResponseWriter) {
	for {
		select {
		case storeResponse := <-srchan:
			// responses received
			log.Printf("INFO: received reply from storage: %v uid: %s", storeResponse, uid)

			//send respose created
			response(w, "Created", http.StatusCreated)
			// remove channel from srchans
			delete(srchans, uid)
			return
		case <-time.After(10 * time.Second):
			// timeout
			log.Printf("ERROR: request timeout uid: %s", uid)

			// send response back to client
			response(w, "Timeout", http.StatusRequestTimeout)

			// remove channel from srchans
			delete(srchans, uid)
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

func response(w http.ResponseWriter, resp string, status int) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, resp)
}
