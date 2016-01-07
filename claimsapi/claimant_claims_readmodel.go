package main

type ClaimantClaims struct {
	ClaimantID string  `json:"claimantId"`
	Claims     []Claim `json:"claims"`
}

type Claim struct {
	Subject string `json:"subject"`
	Claim   string `json:"claim"`
}
