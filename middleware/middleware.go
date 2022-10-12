package middleware

import (
	"assignment-golang-backend/utils"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

//JWT Section
func CheckAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		now := time.Now()

		authCheck := c.Request.Header["Authorization"]
		if len(authCheck) < 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Request failed. Authorization not found.",
			})
			c.Abort()
			return
		}

		//Bearer Token
		authString := authCheck[0]
		tokenCheck := strings.Split(authString, " ")
		if len(tokenCheck) <= 1 {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Request failed. Token not found.",
			})
			c.Abort()
			return
		}

		token := tokenCheck[1]

		email, wallet_id, err := utils.CheckToken(token)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "error",
				"message": "Request failed. Invalid token.",
			})
			c.Abort()
			return
		}

		c.Set("email", email)
		c.Set("wallet_id", wallet_id)

		c.Next()

		log.Println(time.Since(now))
	}
}
