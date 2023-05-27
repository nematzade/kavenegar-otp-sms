package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kavenegar/kavenegar-go"
	"net/http"
)

func sendSms(c *gin.Context) {
	api := kavenegar.New("4F58476A77637158625063507665656B62444241446E447144617750364363496A415352703956415341413D")
	receptor := c.Param("receptor")
	template := "sejamOtp"
	token := "۱۲۳۴"
	params := &kavenegar.VerifyLookupParam{}

	if _, err := api.Verify.Lookup(receptor, template, token, params); err != nil {
		switch err := err.(type) {
		case *kavenegar.APIError:
			c.IndentedJSON(http.StatusOK, gin.H{"status": false, "message": err.Error()})
		case *kavenegar.HTTPError:
			c.IndentedJSON(http.StatusOK, gin.H{"status": false, "message": err.Error()})
		default:
			c.IndentedJSON(http.StatusOK, gin.H{"status": false, "message": err.Error()})
		}
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"status": true})
	}
}

func main() {
	router := gin.Default()
	router.GET("/sendSms/:receptor", sendSms)

	router.Run("localhost:8080")
}
