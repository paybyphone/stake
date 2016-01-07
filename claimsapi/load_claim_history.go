package main

import (
	"encoding/json"
	"strings"

	"github.com/paybyphone/stake"
	"github.com/robertreppel/hist/filestore"
)

func loadClaimHistory(claimantID string) []interface{} {
	// log.Printf("DEBUG loadClaimHistory for '%s:%s'", claimantAggregateType, claimantID)
	if strings.TrimSpace(claimantID) == "" {
		var emptyList []interface{}
		return emptyList
	}
	store, err := filestore.FileStore(dataStoreDirectory)
	failIf(err)

	eventHistory, err := store.Get(claimantAggregateType, claimantID)
	// log.Printf("DEBUG loadClaimHistory - found %d events.", len(eventHistory))
	failIf(err)
	var events []interface{}
	for _, item := range eventHistory {
		if item.Type == "stake.ClaimantCreated" {
			var event stake.ClaimantCreated
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
		if item.Type == "stake.ClaimantDeleted" {
			var event stake.ClaimantDeleted
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
		if item.Type == "stake.ClaimMade" {
			var event stake.ClaimMade
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
		if item.Type == "stake.ClaimRetracted" {
			var event stake.ClaimRetracted
			err := json.Unmarshal(item.Data, &event)
			failIf(err)
			events = append(events, event)
		}
	}
	return events
}
