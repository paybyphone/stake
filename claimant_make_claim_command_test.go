package stake

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMakeClaim(t *testing.T) {
	subjectName := "email"
	claimantID := "jmiller"
	claim := "joe@someemail.com"
	Convey("Given that claimant "+claimantID+" exists and doesn't have any claims yet", t, func() {
		claimant, _, _ := CreateClaimant(claimantID)
		Convey("when a new claim is made", func() {
			response, err := claimant.MakeClaim(subjectName, claim)
			if err != nil {
				panic(err)
			}
			Convey("then the operation succeeds", func() {
				So(response, ShouldEqual, "Success")
			})
			Convey("and two events have occurred:", func() {
				So(len(claimant.Changes), ShouldEqual, 2)
				Convey("The original 'ClaimantCreated',", func() {
					claimantCreated := claimant.Changes[0].(ClaimantCreated)
					So(claimantCreated, ShouldNotBeNil)
				})
				Convey("followed by the new 'ClaimMade' event.", func() {
					claimMade := claimant.Changes[1].(ClaimMade)
					So(claimMade, ShouldNotBeNil)
					Convey("'ClaimMade' should have a claimantID, subject and claim.", func() {
						So(claimMade.ClaimantID, ShouldNotBeBlank)
						So(claimMade.Subject, ShouldNotBeBlank)
						So(claimMade.Claim, ShouldNotBeBlank)
					})
				})
			})
		})
	})

	Convey("When making a claim with a blank subject", t, func() {
		claimant, _, _ := CreateClaimant(claimantID)
		response, err := claimant.MakeClaim("  ", claim)
		if err != nil {
			panic(err)
		}
		numberOfEventsBeforeCommandWasExecuted := 1
		Convey("then the response is 'Subject cannot be blank'", func() {
			So(response, ShouldEqual, "Subject cannot be blank")
		})
		Convey("and no events have resulted.", func() {
			So(len(claimant.Changes), ShouldEqual, numberOfEventsBeforeCommandWasExecuted)
		})
	})

	Convey("When a claimant has a claim", t, func() {
		claimant, _, _ := CreateClaimant(claimantID)
		claimant.MakeClaim(subjectName, claim)
		numberOfEventsBeforeCommandWasExecuted := len(claimant.Changes)
		Convey("and then makes an identical claim", func() {
			response, err := claimant.MakeClaim(subjectName, claim)
			if err != nil {
				panic(err)
			}
			Convey("then the response is 'Claim already exists'", func() {
				So(response, ShouldEqual, "Claim already exists")
			})
			Convey("and no events have resulted.", func() {
				So(len(claimant.Changes), ShouldEqual, numberOfEventsBeforeCommandWasExecuted)
			})
		})
	})
}
