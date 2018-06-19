package main

import (
	"os"
	"path"

	"github.com/urfave/cli"
)

const (
	VERSION = "0.0.8"
	CORS    = "http://genome.compbio.cs.cmu.edu:8080"
	DIR     = ".cnbData"
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
			Usage:  "start an data server",
			Action: CmdStart,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input data tsv/xls/google sheet id",
					Value: "",
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8080,
				},
				cli.StringFlag{
					Name:  "root,r",
					Usage: "root directory",
					Value: home,
				},
			},
		},
		{
			Name:   "rproxy",
			Usage:  "start an data server reverse proxy",
			Action: CmdRP,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input data google sheet id",
					Value: "",
				},
				cli.StringFlag{
					Name:  "title,t",
					Usage: "sheet title",
					Value: "Sheet1",
				},
				cli.IntFlag{
					Name:  "port,p",
					Usage: "data server port",
					Value: 8080,
				},
				cli.StringFlag{
					Name:  "root,r",
					Usage: "root directory",
					Value: home,
				},
			},
		},

		{
			Name:   "maintain",
			Usage:  "maintain data servers",
			Action: CmdMaintain,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "input,i",
					Usage: "input data google sheet id",
					Value: "",
				},
				cli.StringFlag{
					Name:  "title,t",
					Usage: "sheet title",
					Value: "Sheet1",
				},
				cli.StringFlag{
					Name:  "root,r",
					Usage: "root directory",
					Value: home,
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
		{
			Name:   "userdb",
			Usage:  "user db info",
			Action: CmdUserDb,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "i,input", //TODO
					Usage: "input file",
					Value: path.Join(home, ".cnb/user.db"),
				},
			},
		},
	}
	app.Run(os.Args)
}
