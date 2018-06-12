package main

import (
	"github.com/gorilla/mux"
	"github.com/nimezhu/data"
)

func addDataServer(uri string, router *mux.Router, indexRoot string) {
	/*
		genomes := "hg19,mm10,hg38,mm9"
		cytoBandManager := data.NewCytoBandManager("band")
		gs := strings.Split(genomes, ",")
		for _, v := range gs {
			cytoBandManager.Add(v)
		}
		cytoBandManager.ServeTo(router)
	*/
	l := data.NewLoader(indexRoot) //TODO set Index Root
	l.Load(uri, router)
}
