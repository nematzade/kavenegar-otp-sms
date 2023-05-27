package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kavenegar/kavenegar-go"
	"net/http"
)

func sendSms(c *gin.Context) {
	api := kavenegar.New("ApiKey")
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
