package main

import "github.com/paybyphone/stake"

func deleteClaimantsClaimAboutSubject(claimantID string, subject string, claim string) string {
	claimHistory := loadClaimHistory(claimantID)
	if len(claimHistory) == 0 {
		return notFound
	}
	claimant, err := stake.LoadClaimantFromHistory(claimHistory)
	failIf(err)
	if claimant == nil {
		return notFound
	}
	result, err := claimant.RetractClaim(subject, claim)
	failIf(err)
	updateClaimHistory(claimantID, claimant.Changes)

	// Delete from the the list containing everybody who made a claim about the same subject:
	historyOfSubjectOfPreviousClaim := loadClaimHistory(subject)
	claimantsOfSubjectOfDeletedClaim, err := stake.LoadClaimantFromHistory(historyOfSubjectOfPreviousClaim)
	failIf(err)
	claimantsOfSubjectOfDeletedClaim.RetractClaim(claimantID, claim)
	updateClaimHistory(claimantsOfSubjectOfDeletedClaim.ID, claimantsOfSubjectOfDeletedClaim.Changes)

	return result
}
