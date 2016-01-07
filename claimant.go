package stake

import "fmt"

//A Claimant can make claims.
type Claimant struct {
	ID      string
	claims  map[string]bool
	Changes []interface{}
}

func (claimant *Claimant) transition(event interface{}) {
	switch e := event.(type) {
	case ClaimantCreated:
		claimant.ID = e.ClaimantID
		claimant.claims = make(map[string]bool)
	case ClaimMade:
		newClaimKey := claimant.claimKeyFor(e.Subject, e.Claim)
		claimant.claims[newClaimKey] = true
	case ClaimRetracted:
		claimKey := claimant.claimKeyFor(e.Subject, e.Claim)
		delete(claimant.claims, claimKey)
	}
}

func (claimant *Claimant) claimKeyFor(subject string, claim string) string {
	return fmt.Sprintf("%s:%s", subject, claim)
}

// LoadClaimantFromHistory restores the current state of a claimant from past events.
func LoadClaimantFromHistory(events []interface{}) (*Claimant, error) {
	claimant := &Claimant{}
	for _, event := range events {
		claimant.transition(event)
	}
	return claimant, nil
}

func (claimant *Claimant) trackChange(event interface{}) {
	claimant.transition(event)
	claimant.Changes = append(claimant.Changes, event)
}
