package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	go main()
	time.Sleep(2 * time.Second)
}

const claimantID string = "111"
const subject string = "username"
const claim string = "jmiller"

func TestDeleteClaim(t *testing.T) {
	Convey("Given a claim", t, func() {
		postClaim(claimantID, subject, claim)
		Convey("when the claim is deleted", func() {
			status, body := deleteClaim(claimantID, subject, claim)
			Convey("then the status should be 200 OK", func() {
				So(status, ShouldEqual, "200 OK")
			})
			Convey("And the response should be 'Success'", func() {
				So(body, ShouldContainSubstring, "Success")
			})
		})
	})

	Convey("Get a list of everybody who claimed something: If a claim was retracted, it no longer appears in the list.", t, func() {
		_, body := getClaims(subject)
		Convey("the response should contain the previous claim's claimantID as the subject: "+claimantID, func() {
			So(body, ShouldNotContainSubstring, "\"subject\":\""+claimantID+"\"")
		})
		Convey("and the the previous claim '"+claim+"' as the claim.", func() {
			So(body, ShouldNotContainSubstring, "\"claim\":\""+claim+"\"")
		})
	})
}

func deleteClaim(claimantID string, subject string, claim string) (string, string) {

	url := fmt.Sprintf("http://localhost:8090/claimants/%s/claims", claimantID)
	payload := strings.NewReader(fmt.Sprintf("{ \"subject\" : \"%s\", \"claim\" : \"%s\"}", subject, claim))
	req, _ := http.NewRequest("DELETE", url, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	return res.Status, string(body)
}
