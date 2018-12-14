package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type agencyID struct {
	AgencyID     int
	AgencyNumber string
}

func main() {
	filename := "c:\\Users\\tilliad\\OneDrive - TUI\\#dt\\agendaagency\\agenda_agenturen_mini.txt"
	filenameOut := "c:\\Users\\tilliad\\OneDrive - TUI\\#dt\\agendaagency\\agenda_agenturen_json.txt"

	agencyids := read(filename)
	jsons := convert(agencyids)
	write(jsons, filenameOut)
}

func read(filename string) []agencyID {
	f, _ := os.Open(filename)

	reader := csv.NewReader(f)
	reader.Comma = ';'
	var agencyids []agencyID
	for {
		record, err := reader.Read()
		// Stop at EOF.
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		agencyids = append(agencyids, agencyID{
			AgencyID:     0,
			AgencyNumber: record[0]})
	}
	return agencyids
}

func convert(agencyIds []agencyID) []string {
	var jsons []string
	for _, id := range agencyIds {
		jsonObj, _ := json.Marshal(id)
		jsons = append(jsons, string(jsonObj))
	}
	return jsons
}

func write(jsons []string, filename string) error {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, jsonObj := range jsons {
		fmt.Fprintln(w, jsonObj)
	}
	return w.Flush()
}
