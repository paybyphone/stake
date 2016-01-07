package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"os"
	"reflect"
)

type eventlogEntry struct {
	EventType  string
	ClaimantID string
	Subject    string
	Claim      string
}

func logEvents(events []interface{}, logFilePath string) {
	f, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()
	w := csv.NewWriter(f)
	for _, e := range events {
		logEntryBytes, err := json.Marshal(e)
		if err != nil {
			log.Fatal(err)
		}
		var logEntry eventlogEntry
		err = json.Unmarshal(logEntryBytes, &logEntry)
		if err != nil {
			log.Fatal(err)
		}
		logEntry.EventType = reflect.TypeOf(e).String()
		logEntryStr := make([]string, 5)
		logEntryStr[0] = logEntry.EventType
		logEntryStr[1] = logEntry.ClaimantID
		logEntryStr[2] = logEntry.Subject
		logEntryStr[3] = logEntry.Claim
		if err := w.Write(logEntryStr); err != nil {
			log.Fatal(err)
		}
	}
	w.Flush()
	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
