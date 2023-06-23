package main

import (
	"net/http"
	"strconv"
	"time"
)

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
