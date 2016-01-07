package stake

import (
	"strings"
)

// RetractClaim deletes a previously made claim
func (claimant *Claimant) RetractClaim(subject string, claim string) (string, error) {
	if strings.TrimSpace(subject) == "" {
		return "Subject cannot be blank", nil
	}

	claimExists := claimant.claimKeyFor(subject, claim)
	if claimant.claims[claimExists] == false {
		return "Claim not found", nil
	}

	claimant.trackChange(ClaimRetracted{ClaimantID: claimant.ID, Subject: subject, Claim: claim})
	return "Success", nil
}
