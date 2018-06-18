package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
	"time"

	"github.com/fatih/color"
	"github.com/nimezhu/asheets"
	"github.com/nimezhu/data"
	"github.com/olekukonko/tablewriter"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli"
	sheets "google.golang.org/api/sheets/v4"
)

func CmdMaintain(c *cli.Context) {
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
	v, _ := asheets.ReadSheet(title, srv, sheetid, "A")
	head, rowid, valueMap := v.ColIds, v.RowIds, v.RowValue
	name2col := make(map[string]string)
	name2idx := make(map[string]int)
	for i, h := range head {
		name2col[strings.ToLower(h)] = asheets.NumberToColName(i + 1)
		name2idx[strings.ToLower(h)] = i
	}
	values := [][]interface{}{}
	data := [][]string{}
	repairs := []string{}
	generateScript := func(rid string) string {
		cmd := valueMap[rid][name2idx["cmd"]]
		port := valueMap[rid][name2idx["port"]]
		sheetid := valueMap[rid][name2idx["sheetid"]]
		s := fmt.Sprintf("nohup %s start -i %s -p %s > %s.log& 2>&1", cmd, sheetid, port, port)
		return s
	}
	for _, rid := range rowid {
		url := valueMap[rid][name2idx["uri"]]
		res, err2 := http.Get(url + "/version")
		dat := time.Now()
		if err2 != nil {
			values = append(values, []interface{}{"not active", dat.String(), "Null"})
			data = append(data, []string{color.RedString(rid), url, color.RedString("Not Active")})
			if valueMap[rid][name2idx["cmd"]] == "cnbData" && GuessURIType(valueMap[rid][name2idx["sheetid"]]) == "gsheet" {
				repairs = append(repairs, rid)
			}
		} else {
			body, err3 := ioutil.ReadAll(res.Body)
			if err3 == nil {
				app := gjson.Get(string(body), "Appname")
				version := gjson.Get(string(body), "Version")
				data = append(data, []string{rid, url, app.String() + " " + version.String()})
				values = append(values, []interface{}{"active", dat.String(), app.String() + " " + version.String()})
			} else {
				values = append(values, []interface{}{"not recogonize", dat.String(), "Unknown"})
			}
		}

	}

	rangeData := title + "!" + name2col["status"] + "2"
	//values := [][]interface{}{{"sample_A1", "sample_C1"}, {"sample_A2", "sample_C2"}, {"sample_A3", "sample_A3"}}
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
	fmt.Println("Status:")
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "URI", "Status"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render() // Send output
	fmt.Println("Suggest Scripts:")
	//Suggesting Script
	for _, v := range repairs {
		fmt.Println(generateScript(v))
	}

}
