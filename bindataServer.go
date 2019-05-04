package main

//go:generate go-bindata-assetfs -pkg main app/... client_secret.json
import (
	"net/http"
)

func BindataServer(root string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := root + r.RequestURI
		bytes, err := Asset(id)
		if err != nil {
			w.Write([]byte("file not found"))
		} else {
			w.Write(bytes)
		}

	})
}
