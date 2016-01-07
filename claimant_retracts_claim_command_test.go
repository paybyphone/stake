package stake

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRetractNeverMadeClaim(t *testing.T) {
	subjectName := "administrator"
	claimantID := "jmiller"
	claim := "true"
	Convey("When a claimant retracts a claim she/he never made", t, func() {
		claimant, _, _ := CreateClaimant(claimantID)
		numberOfEventsBeforeCommandWasExecuted := len(claimant.Changes)
		response, err := claimant.RetractClaim(subjectName, claim)
		if err != nil {
			panic(err)
		}
		Convey("then the response is Not Found", func() {
			So(response, ShouldEqual, "Claim not found")
		})
		Convey("and no events have resulted.", func() {
			So(len(claimant.Changes), ShouldEqual, numberOfEventsBeforeCommandWasExecuted)
		})
	})
}

func TestRetractPreviouslyMadeClaim(t *testing.T) {
	subjectName := "administrator"
	claimantID := "jmiller"
	claim := "true"
	Convey("When a claimant retracts a claim she/he never made", t, func() {
		claimant, _, _ := CreateClaimant(claimantID)
		claimant.MakeClaim(subjectName, claim)
		numberOfEventsBeforeCommandWasExecuted := len(claimant.Changes)
		response, err := claimant.RetractClaim(subjectName, claim)
		if err != nil {
			panic(err)
		}
		Convey("then the response is 'Sucess'", func() {
			So(response, ShouldEqual, "Success")
		})
		Convey("and a ClaimRetracted event has occurred.", func() {
			So(len(claimant.Changes), ShouldEqual, numberOfEventsBeforeCommandWasExecuted+1)
			thirdEvent := claimant.Changes[2].(ClaimRetracted)
			So(thirdEvent, ShouldNotBeNil)
		})
		Convey("and the claimant no longer has the claim.", func() {
			claimKey := claimant.claimKeyFor(subjectName, claim)
			hasClaim := claimant.claims[claimKey]
			So(hasClaim, ShouldBeFalse)
		})
	})
}
