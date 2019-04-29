package main

//go:generate go-bindata-assetfs -pkg main web/...
import (
	"log"
	"net/http"
)

func BindataServer() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		l := len(r.RequestURI)
		id := r.RequestURI[1:l]
		log.Println(id)
		if id == "web" || id == "web/" {
			id = "web/index.html"
		}
		bytes, err := Asset(id)
		if err != nil {
			w.Write([]byte("file not found"))
		} else {
			w.Write(bytes)
		}

	})
}
