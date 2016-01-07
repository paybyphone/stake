package main

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/paybyphone/stake"
)

const dataStoreDirectory string = "/tmp/stake-eventstore"
const logFilePath = "/tmp/stake-eventlog.csv"

func main() {
	r := gin.Default()
	r.POST("/claimants/:id/claims", handleClaimantsClaimsPost)
	r.GET("/claimants/:id/claims", handleClaimantsClaimsGet)
	r.DELETE("/claimants/:id/claims", handleClaimantsClaimsDelete)
	r.Run(":8090")
}

func handleClaimantsClaimsPost(c *gin.Context) {
	claimantID := c.Param("id")
	if strings.TrimSpace(claimantID) == "" {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	var claim stake.Claim
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&claim)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	if strings.TrimSpace(claim.Claim) == "" || strings.TrimSpace(claim.Subject) == "" {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	result := claimantMakesClaim(claimantID, claim.Subject, claim.Claim)
	if result == success {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusBadRequest, result)
	}
}

func handleClaimantsClaimsGet(c *gin.Context) {
	claimantID := c.Param("id")

	claimantClaims, result := getClaimsMadeBy(claimantID)
	if result == success {
		c.JSON(http.StatusOK, claimantClaims)
	} else if result == notFound {
		c.JSON(http.StatusNotFound, result)
	}
}

func handleClaimantsClaimsDelete(c *gin.Context) {
	claimantID := c.Param("id")
	if strings.TrimSpace(claimantID) == "" {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	var claim stake.Claim
	decoder := json.NewDecoder(c.Request.Body)
	err := decoder.Decode(&claim)
	if err != nil {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	if strings.TrimSpace(claim.Claim) == "" || strings.TrimSpace(claim.Subject) == "" {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	if strings.TrimSpace(claimantID) == "" || strings.TrimSpace(claim.Claim) == "" || strings.TrimSpace(claim.Subject) == "" {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}
	result := deleteClaimantsClaimAboutSubject(claimantID, claim.Subject, claim.Claim)
	if result == success {
		c.JSON(http.StatusOK, result)
	} else {
		c.JSON(http.StatusNotFound, result)
	}
}

const success string = "Success"
const notFound string = "Not Found"
