package main

import (
	"os"

	"github.com/rs/cors"
	"github.com/urfave/cli"
)

const (
	VERSION = "0.0.10"
	DIR     = ".cnbData"
)

var (
	CORS        = []string{"http://genome.compbio.cs.cmu.edu:8080", "http://x7.andrew.cmu.edu:8080", "chrome-extension://djcdicpaejhpgncicoglfckiappkoeof", "chrome-extension://gedcoafficobgcagmjpplpnmempkcpfp", "https://genome.compbio.cs.cmu.edu"}
	corsOptions = cors.Options{
		AllowedOrigins:   CORS,
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization"},
	}
)

func main() {
	app := cli.NewApp()
	app.Version = VERSION
	app.Name = "cnb dataserver tools"
	app.Usage = "cnbData start -i [[google_sheet_id OR xls file]] -p [[port]]"
	app.EnableBashCompletion = true //TODO
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "verbose",
			Usage: "Show more output",
		},
	}
	home := os.Getenv("HOME")
	app.Commands = []cli.Command{
		{
			Name:   "start",
			Usage:  "start a data server",
			Action: CmdStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input data xls/google sheet id",
					Value: "",
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8080,
				},
				cli.StringFlag{
					Name:  "root,r",
					Usage: "root directory, where store credentials and index files, defualt is $HOME/.cnbData",
					Value: home,
				},
				cli.BoolFlag{
					Name:  "local,l",
					Usage: "serve 127.0.0.1 only",
				},
				cli.StringFlag{
					Name:  "code,c",
					Usage: "set password for server, override -l",
					Value: "",
				},
			},
		},
		{
			Name:   "file",
			Usage:  "start a file server",
			Action: CmdFile,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "root,r", //TODO
					Usage: "root directory",
					Value: home,
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8080,
				},
			},
		},
	}
	app.Run(os.Args)
}
