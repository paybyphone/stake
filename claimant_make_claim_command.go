package stake

import "strings"

// MakeClaim lets a claimant make a claim about a subject.
func (claimant *Claimant) MakeClaim(subject string, claim string) (string, error) {
	if strings.TrimSpace(subject) == "" {
		return "Subject cannot be blank", nil
	}

	claimAlreadyExists := claimant.claimKeyFor(subject, claim)
	if claimant.claims[claimAlreadyExists] == true {
		return "Claim already exists", nil
	}

	claimant.trackChange(ClaimMade{ClaimantID: claimant.ID, Subject: subject, Claim: claim})
	return "Success", nil
}
