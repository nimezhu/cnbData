package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func addTmplBindata(router *mux.Router, para interface{}) {
	router.HandleFunc("/app/{page}.html", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Accept-Encoding")
		w.Header().Add("Content-Type", "text/html")
		ps := mux.Vars(r)
		bytes, _ := Asset("app/html/" + ps["page"] + ".tmpl")
		tmpl := template.New("html")
		tmpl, err := tmpl.Parse(string(bytes))
		dir, _ := AssetDir("app/module")
		for _, d := range dir {
			bytes, err1 := Asset("app/module/" + d)
			if err1 != nil {
				log.Panicf("Unable to parse: template=%s, err=%s", d, err)
			}
			tmpl.New(d).Parse(string(bytes))
		}
		if err != nil {
			log.Println("error parse template")
		}
		err = tmpl.Execute(w, para) //constant
		if err != nil {
			log.Println("error executing template")
		}
	})
}
