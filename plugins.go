package main

import (
	"errors"
	"fmt"

	"github.com/nimezhu/data"
	"github.com/nimezhu/tbl2x"
)

/* TODO : formalize to tsv ome Browser
 *        for future data loader
 */
func pluginTsv(dbname string, data interface{}) (data.DataRouter, error) {
	switch v := data.(type) {
	default:
		fmt.Printf("unexpected type %T", v)
		return nil, errors.New(fmt.Sprintf("bigwig format not support type %T", v))
	case string:
		return nil, errors.New("todo")
	case map[string]interface{}:
		r := &tbl2x.TableRouter{dbname, make(map[string]*tbl2x.Table)}
		err := r.Load(data.(map[string]interface{}))
		return r, err
	}
}
