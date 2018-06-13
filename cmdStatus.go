package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strings"
	"time"

	"github.com/nimezhu/asheets"
	"github.com/nimezhu/data"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
	sheets "google.golang.org/api/sheets/v4"
)

func CmdStatus(c *cli.Context) {
	dir := path.Join(c.String("root"), DIR)
	ctx := context.Background()
	title := c.String("title")
	sheetid := c.String("input")
	config := data.GsheetConfig()
	gA := asheets.NewGAgent(dir)
	client := gA.GetClient(ctx, config)
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets Client %v", err)
	}
	head, rowid, valueMap := asheets.ReadSheet(title, srv, sheetid, "A")
	//fmt.Println(valueMap)
	for i, r := range rowid {
		fmt.Println(i, r)
	}
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
	values := [][]interface{}{}
	for _, rid := range rowid {
		url := valueMap[rid][uriColIdx]
		res, err2 := http.Get(url + "/version")
		dat := time.Now()
		if err2 != nil {
			values = append(values, []interface{}{"not active", dat.String(), "Null"})
		} else {
			body, err3 := ioutil.ReadAll(res.Body)
			if err3 == nil {
				app := gjson.Get(string(body), "Appname")
				version := gjson.Get(string(body), "Version")
				fmt.Println(app.String(), version.String())
				values = append(values, []interface{}{"active", dat.String(), app.String() + " " + version.String()})
			} else {
				values = append(values, []interface{}{"not recogonize", dat.String(), "Unknown"})
			}
		}

	}

	//TODO make it into function
	rangeData := title + "!" + statusCol + "2"
	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "RAW",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  rangeData,
		Values: values,
	})

	_, err4 := srv.Spreadsheets.Values.BatchUpdate(sheetid, rb).Context(ctx).Do()
	if err4 != nil {
		log.Fatalf("Unable to write data to sheet. %v", err4)
	}
	fmt.Println("Done")

}
