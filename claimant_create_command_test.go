package stake

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCreateClaimant(t *testing.T) {
	claimantID := "2778c36f-c35c-4254-8f39-2f46308869cc"
	Convey("Given that the member with id  "+claimantID+" is not a claimant", t, func() {
		Convey("when we try to create the claimant", func() {
			claimant, response, err := CreateClaimant(claimantID)
			if err != nil {
				panic(err)
			}
			Convey("then the response is 'Success'", func() {
				So(response, ShouldEqual, "Success")
			})
			Convey("and event is published", func() {
				So(len(claimant.Changes), ShouldEqual, 1)
			})
			Convey("and the event type is 'ClaimantCreated'.", func() {
				claimantCreated := claimant.Changes[0].(ClaimantCreated)
				So(claimantCreated, ShouldNotBeNil)
			})
			Convey("and the claimant has the ID "+claimantID, func() {
				So(claimant.ID, ShouldEqual, claimantID)
			})
		})
	})
	Convey("Given that the claimant ID is blank", t, func() {
		Convey("when we try to create the claimant", func() {
			claimant, response, err := CreateClaimant("")
			if err != nil {
				panic(err)
			}
			Convey("then no claimant is created", func() {
				So(claimant, ShouldBeNil)
			})
			Convey("and the response is 'ID cannot be blank'", func() {
				So(response, ShouldEqual, "ID cannot be blank")
			})
		})
	})
}
