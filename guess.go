package main

import (
	"os"
	"path"
	"regexp"
)

func GuessURIType(uri string) string {
	http, _ := regexp.Compile("^http://")
	https, _ := regexp.Compile("^https://")
	if len(uri) == 44 {
		if _, err := os.Stat(uri); os.IsNotExist(err) {
			if !http.MatchString(uri) && !https.MatchString(uri) {
				return "gsheet"
			}
		}
	}
	ext := path.Ext(uri)
	if ext == ".json" {
		return "json" //not for obsoleted version
	} else if ext == ".xlsx" || ext == "xls" {
		return "xls"
	} else if ext == "txt" {
		return "txt"
	}
	return "unknown"
}
