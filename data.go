package main

import "net/http"

type DataIndex struct {
	Genome string      `json:"genome"`
	Dbname string      `json:"dbname"`
	Data   interface{} `json:"data"` // map[string]string or map[string][]string? could be uri or more sofisticated data structure such as binindex image
	Format string      `json:"format"`
}

func cred(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Authorization")
		next.ServeHTTP(w, r)
	})
}

type App struct {
	Appname string
	Version string
}
