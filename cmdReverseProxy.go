package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/nimezhu/asheets"
	"github.com/nimezhu/data"
	"github.com/urfave/cli"
	sheets "google.golang.org/api/sheets/v4"
)

type App struct {
	Appname string
	Version string
}

var rpApp = App{
	"CMU Reverse Proxy Dataome Browser",
	"0.0.1",
}

func CmdRP(c *cli.Context) {
	ctx := context.Background()
	root := c.String("root")
	dir := path.Join(root, DIR)
	title := c.String("title")
	sheetid := c.String("input")
	port := c.Int("port")
	config := data.GsheetConfig()
	gA := asheets.NewGAgent(dir)
	client := gA.GetClient(ctx, config)
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}

	//Wrapper to Class
	head, rowid, valueMap := asheets.ReadSheet(title, srv, sheetid, "A")
	//fmt.Println(valueMap)
	for i, r := range rowid {
		fmt.Println(i, r)
	}
	//TODO Wrapper to Class
	name2idx := make(map[string]int)
	for i, h := range head {
		name2idx[strings.ToLower(h)] = i
	}
	uriColIdx := name2idx["uri"]
	var genomes []string
	genomeMap := make(map[string][]string)
	for _, rid := range rowid {
		url := valueMap[rid][uriColIdx]
		res, err2 := http.Get(url + "/genomes")
		if err2 == nil {
			body, err3 := ioutil.ReadAll(res.Body)
			if err3 == nil {
				fmt.Println(rid, string(body))
				json.Unmarshal(body, &genomes)
				for i, g := range genomes {
					fmt.Println(i, g)
					if v, ok := genomeMap[g]; !ok {
						genomeMap[g] = []string{rid}
					} else {
						genomeMap[g] = append(v, rid)
					}
				}
			}
		}
	}

	/*TODO Functionalize it */
	genomes = []string{}
	for g, v := range genomeMap {
		fmt.Println(g, v)
		genomes = append(genomes, g)
		//add genome handler map ...
	}
	// merge genome urls

	fmt.Println("Done")
	router := mux.NewRouter()
	cors := data.CorsFactory("*")
	router.Use(cors)
	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		a, _ := json.Marshal(rpApp)
		w.Write(a)
	})
	router.HandleFunc("/genomes", func(w http.ResponseWriter, r *http.Request) {
		a, _ := json.Marshal(genomes)
		w.Write(a)
	})
	router.HandleFunc("/{genome}/ls", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		g := vars["genome"]
		dbAll := []DataIndex{}
		for _, rid := range genomeMap[g] {
			uri := valueMap[rid][uriColIdx]
			var dbs []DataIndex
			res, err4 := http.Get(uri + "/" + g + "/ls")
			if err4 == nil {
				body, err5 := ioutil.ReadAll(res.Body)
				if err5 == nil {
					json.Unmarshal(body, &dbs)
					for _, v := range dbs {
						dbAll = append(dbAll, v)
					}
				}
			}
		}
		a, _ := json.Marshal(dbAll)
		w.Write(a)
	})

	dbnameServerMap := map[string]map[string]string{}
	for _, g := range genomes {
		for _, rid := range genomeMap[g] {
			uri := valueMap[rid][uriColIdx]
			var dbs []DataIndex
			res, err6 := http.Get(uri + "/" + g + "/ls")
			if err6 == nil {
				body, err7 := ioutil.ReadAll(res.Body)
				if err7 == nil {
					json.Unmarshal(body, &dbs)
					for _, v := range dbs {
						if _, ok := dbnameServerMap[g]; !ok {
							dbnameServerMap[g] = make(map[string]string)
						}
						dbnameServerMap[g][v.Dbname] = uri
					}
				}
			}
		}
	}

	router.HandleFunc("/{genome}/{dbname}/list", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		g := vars["genome"]
		db := vars["dbname"]
		uri := dbnameServerMap[g][db] + "/" + g + "/" + db + "/" + "list?attr=1"
		res, err8 := http.Get(uri)
		if err8 == nil {
			body, err9 := ioutil.ReadAll(res.Body)
			if err9 == nil {

				w.Write(body)
			}
		}
		//w.Write([]byte(uri))
	})
	router.HandleFunc("/{genome}/{dbname}/{cmd:.*}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		g := vars["genome"]
		db := vars["dbname"]
		cmd := vars["cmd"]
		uri := dbnameServerMap[g][db] + "/" + g + "/" + db + "/" + cmd
		res, err8 := http.Get(uri)
		if err8 == nil {
			body, err9 := ioutil.ReadAll(res.Body)
			if err9 == nil {

				w.Write(body)
			}
		}
	})
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router}
	log.Println("Please open http://127.0.0.1:" + strconv.Itoa(port))
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
