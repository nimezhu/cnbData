package main

import (
	"context"
	"log"
	"path"

	"github.com/gorilla/mux"
	"github.com/nimezhu/asheets"
	"github.com/nimezhu/data"
	"github.com/nimezhu/sand"
	"github.com/urfave/cli"
	"golang.org/x/oauth2/google"
)

func CmdStart(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	mode := "w"
	root := c.String("root")
	router := mux.NewRouter()

	dir := path.Join(root, DIR)
	ctx := context.Background()
	b, err := data.Asset("client_secret.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/auth/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	gA := asheets.NewGAgent(dir)
	if !gA.HasCacheFile() {
		gA.GetClient(ctx, config)
	}

	//cred := c.String("cred")
	s := sand.Sand{
		"CMU Dataome Browser",
		root,
		DIR,
		VERSION,
		[]string{},
		[]string{},
		[]string{},
		"S.render",
		make(map[string]string),
		"",
	}
	idxRoot := s.InitIdxRoot(root) //???
	//sand.InitCred(cred)
	addDataServer(uri, router, idxRoot) //TODO
	s.InitRouter(router)
	s.InitHome(root)
	s.Start(mode, port, router)

	return nil
}
