package main

import (
	"log"

	"github.com/paybyphone/stake"
)

func claimantMakesClaim(claimantID string, subjectName string, claim string) string {
	// Record claim:
	claimHistory := loadClaimHistory(claimantID)
	// log.Printf("DEBUG claimsapi.claimantMakesClaim - claimHistory: %v\n", claimHistory)
	var claimant *stake.Claimant
	var err error
	if len(claimHistory) == 0 {
		claimant, _, err = stake.CreateClaimant(claimantID)
		failIf(err)
		updateClaimHistory(claimantID, claimant.Changes)
	} else {
		log.Printf("DEBUG claimsapi.claimantMakesClaim - claimant: %v\n", claimant)
		claimHistory = loadClaimHistory(claimantID)
		failIf(err)
		claimant, err = stake.LoadClaimantFromHistory(claimHistory)
		failIf(err)
	}
	response, err := claimant.MakeClaim(subjectName, claim)
	failIf(err)
	updateClaimHistory(claimantID, claimant.Changes)

	// Record a new claimant for the subject:
	subjectClaimHistory := loadClaimHistory(subjectName)
	log.Printf("DEBUG claimsapi.claimantMakesClaim - CsubjectClaimHistory: %v\n", subjectClaimHistory)

	var subjectTurnedClaimant *stake.Claimant
	if len(subjectClaimHistory) == 0 {
		subjectTurnedClaimant, _, err = stake.CreateClaimant(subjectName)
		updateClaimHistory(subjectTurnedClaimant.ID, subjectTurnedClaimant.Changes)
		log.Printf("DEBUG claimsapi.claimantMakesClaim - Created new claimant: %v\n", subjectTurnedClaimant)
		failIf(err)
	} else {
		subjectClaimHistory = loadClaimHistory(subjectName)
		failIf(err)
		subjectTurnedClaimant, err = stake.LoadClaimantFromHistory(subjectClaimHistory)
		failIf(err)
	}
	response, err = subjectTurnedClaimant.MakeClaim(claimantID, claim)
	failIf(err)

	log.Printf("DEBUG claimsapi.claimantMakesClaim - subjectTurnedClaimant: %v\n", subjectTurnedClaimant)
	updateClaimHistory(subjectTurnedClaimant.ID, subjectTurnedClaimant.Changes)

	return response
}
