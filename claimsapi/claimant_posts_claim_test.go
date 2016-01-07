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

func TestPostClaim(t *testing.T) {
	claimantID := "111"
	subject := "username"
	claim := "jmiller"
	Convey("Given a POST", t, func() {
		status, body := postClaim(claimantID, subject, claim)
		Convey("The status is 200 OK", func() {
			So(status, ShouldEqual, "200 OK")
			Convey("And the response should be 'Success'", func() {
				So(body, ShouldContainSubstring, "Success")
			})
		})
	})

	Convey("Given a GET claims", t, func() {
		status, body := getClaims(claimantID)
		Convey("the status is 200 OK", func() {
			So(status, ShouldEqual, "200 OK")
			Convey("and the response should contain the subject", func() {
				So(body, ShouldContainSubstring, "\"subject\":\"username\"")
			})
			Convey("and the response should contain the claim", func() {
				So(body, ShouldContainSubstring, "\"claim\":\"jmiller\"")
			})
		})
	})

	Convey("Given that the claimant was added to the the list containing everybody who made a claim about the same subject", t, func() {
		_, body := getClaims(subject)
		Convey("when a GET with the subject as the claimantID is made, the response should contain the previous claim's claimantID as the subject", func() {
			So(body, ShouldContainSubstring, "\"subject\":\""+claimantID+"\"")
		})
		Convey("and the the previous claim '"+claim+"' as the claim.", func() {
			So(body, ShouldContainSubstring, "\"claim\":\""+claim+"\"")
		})
	})
}

func TestPostClaimWithBlankParameters(t *testing.T) {
	claimantID := ""
	subject := "username"
	claim := "jmiller"
	Convey("Given a POST", t, func() {
		status, body := postClaim(claimantID, subject, claim)
		Convey("The status is 400 Bad Request", func() {
			So(status, ShouldEqual, "400 Bad Request")
			Convey("And the response should be '400 Bad Request'", func() {
				So(body, ShouldContainSubstring, "Bad Request")
			})
		})
	})
}

func postClaim(claimantID string, subject string, claim string) (string, string) {
	url := fmt.Sprintf("http://localhost:8090/claimants/%s/claims", claimantID)
	payload := strings.NewReader(fmt.Sprintf("{ \"subject\" : \"%s\", \"claim\" : \"%s\"}", subject, claim))
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	return res.Status, string(body)
}

func getClaims(claimantID string) (string, string) {
	url := fmt.Sprintf("http://localhost:8090/claimants/%s/claims", claimantID)
	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("content-type", "application/json")
	req.Header.Add("cache-control", "no-cache")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	return res.Status, string(body)
}
