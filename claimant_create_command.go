package stake

// CreateClaimant - The claimantID must be unique.
func CreateClaimant(claimantID string) (*Claimant, string, error) {
	if claimantID == "" {
		return nil, "ID cannot be blank", nil
	}
	claimant := Claimant{}
	claimant.trackChange(ClaimantCreated{ClaimantID: claimantID})
	return &claimant, "Success", nil
}
