package main

import "github.com/paybyphone/stake"

func getClaimsMadeBy(claimantID string) (*ClaimantClaims, string) {
	claimHistory := loadClaimHistory(claimantID)
	if len(claimHistory) == 0 {
		return nil, notFound
	}

	claims := make(map[string]Claim)

	for _, claimEvent := range claimHistory {
		switch e := claimEvent.(type) {
		case stake.ClaimMade:
			claims[e.Subject] = Claim{Subject: e.Subject, Claim: e.Claim}
		case stake.ClaimRetracted:
			delete(claims, e.Subject)
		}
	}

	var claimsList []Claim
	for _, claim := range claims {
		claimsList = append(claimsList, claim)
	}

	var claimantClaims ClaimantClaims
	claimantClaims.ClaimantID = claimantID
	claimantClaims.Claims = claimsList

	return &claimantClaims, success
}
