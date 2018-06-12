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
	"golang.org/x/oauth2/google"
	sheets "google.golang.org/api/sheets/v4"
)

type App struct {
	Appname string
	Version string
}

var rpApp = App{
	"Reverse Proxy Dataome Browser",
	"0.0.1",
}

func CmdRP(c *cli.Context) {
	ctx := context.Background()
	root := c.String("root")
	dir := path.Join(root, DIR)
	title := c.String("title")
	sheetid := c.String("input")
	port := c.Int("port")
	b, err := data.Asset("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
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
	var statusCol string
	var lastCheckCol string
	var uriCol string
	for i, h := range head {
		if strings.ToLower(h) == "status" {
			statusCol = asheets.NumberToColName(i + 1)
		}
		if strings.ToLower(h) == "lastchecktime" {
			lastCheckCol = asheets.NumberToColName(i + 1)
		}
		if strings.ToLower(h) == "status" {
			statusCol = asheets.NumberToColName(i + 1)
		}
		if strings.ToLower(h) == "uri" {
			uriCol = asheets.NumberToColName(i + 1)
		}
		fmt.Println(asheets.NumberToColName(i+1), h)
	}
	uriColIdx := asheets.ColNameToNumber(uriCol) - 1
	fmt.Println("Status Col:", statusCol)
	fmt.Println("LastCheckTime Col:", lastCheckCol)
	//values := [][]interface{}{}
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
	router.HandleFunc("/version", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		a, _ := json.Marshal(rpApp)
		w.Write(a)
	})
	router.HandleFunc("/genomes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		a, _ := json.Marshal(genomes)
		w.Write(a)
	})
	router.HandleFunc("/{genome}/ls", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		w.Header().Set("Access-Control-Allow-Origin", "*")
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
				w.Header().Set("Access-Control-Allow-Origin", "*")
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
				w.Header().Set("Access-Control-Allow-Origin", "*")
				w.Write(body)
			}
		}
	})
	server := &http.Server{Addr: ":" + strconv.Itoa(port), Handler: router}
	err = server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	log.Println("Please open http://127.0.0.1:" + strconv.Itoa(port))

}
