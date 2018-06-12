package main

type DataIndex struct {
	Genome string      `json:"genome"`
	Dbname string      `json:"dbname"`
	Data   interface{} `json:"data"` // map[string]string or map[string][]string? could be uri or more sofisticated data structure such as binindex image
	Format string      `json:"format"`
}
