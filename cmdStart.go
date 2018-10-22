package main

import (
	"context"
	"net/http"
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
	root := c.String("root")
	router := mux.NewRouter()
	if GuessURIType(uri) == "gsheet" {
		dir := path.Join(root, DIR)
		ctx := context.Background()
		config := data.GsheetConfig()
		gA := asheets.NewGAgent(dir)
		if !gA.HasCacheFile() {
			gA.GetClient(ctx, config)
		}
	}
	//cred := c.String("cred")
	s := box.Box{
		"CMU Dataome Browser",
		root,
		DIR,
		VERSION,
	}
	s.InitRouter(router)
	s.InitHome(root)
	idxRoot := s.InitIdxRoot(root) //???
	l := data.NewLoader(idxRoot)
	l.Plugins["tsv"] = pluginTsv
	l.Load(uri, router)

	router.Use(cred)
	//router.Use(userMiddleware)
	/* Add User Control
	 * For Specific Group User Email
	 */
	s.StartDataServer(port, router, &corsOptions)

	return nil
}

func strictCorsFactory(sites string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("Origin") == sites {
				next.ServeHTTP(w, r)
			} else {
				w.Write([]byte("not authorized"))
			}
		})
	}
}
