package stake

//A ClaimantCreated event is raised when a new claimant has been created.
type ClaimantCreated struct {
	ClaimantID string
}

// ClaimMade happens when a claimant has made a claim regarding a subject.
// e.g. Joe has claimed that 'joe@somemail.com' is his email or Clara claims that it's true that she is a BackofficeAdmin.
type ClaimMade struct {
	ClaimantID string
	Subject    string
	Claim      string
}

// ClaimRetracted happens when a claim about a subject is no longer being made by the claimant.
type ClaimRetracted struct {
	ClaimantID string
	Subject    string
	Claim      string
}

// ClaimantDeleted events occur when a claimant is deleted.
type ClaimantDeleted struct {
	ClaimantID string
}
