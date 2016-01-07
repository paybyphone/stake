package main

import (
	"encoding/json"
	"reflect"

	"github.com/robertreppel/hist/filestore"
)

func updateClaimHistory(claimantID string, changes []interface{}) {
	store, err := filestore.FileStore(dataStoreDirectory)
	failIf(err)
	for _, event := range changes {
		jsonEvent, err := json.Marshal(event)
		failIf(err)
		store.Save(claimantAggregateType, claimantID, reflect.TypeOf(event).String(), []byte(jsonEvent))
	}
	logEvents(changes, logFilePath)
}

func failIf(err error) {
	if err != nil {
		panic(err)
	}
}

const claimantAggregateType string = "Claimant"
