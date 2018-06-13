package main

import (
	"context"
	"path"

	"github.com/gorilla/mux"
	"github.com/nimezhu/asheets"
	"github.com/nimezhu/box"
	"github.com/nimezhu/data"
	"github.com/urfave/cli"
)

func CmdStart(c *cli.Context) error {
	uri := c.String("input")
	port := c.Int("port")
	mode := "w"
	root := c.String("root")
	router := mux.NewRouter()
	if GuessURIType(uri) == "gsheet" {
		dir := path.Join(root, DIR)
		ctx := context.Background()
		config := defaultConfig()
		gA := asheets.NewGAgent(dir)
		if !gA.HasCacheFile() {
			gA.GetClient(ctx, config)
		}
	}
	//cred := c.String("cred")
	s := box.Box{
		"CMU Dataome Server",
		root,
		DIR,
		VERSION,
	}
	s.InitRouter(router)
	s.InitHome(root)
	idxRoot := s.InitIdxRoot(root) //???
	l := data.NewLoader(idxRoot)
	l.Load(uri, router)
	router.Use(data.CorsMiddleware)
	s.Start(mode, port, router)

	return nil
}
