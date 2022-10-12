package middleware

import (
	"assignment-golang-backend/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(c *gin.Context) {
	authorizationHeader := c.GetHeader("Authorization")
	tokenStr, err := utils.ParseAuthorizationHeader(authorizationHeader)
	if err != nil {
		utils.WriteErrorResponse(
			c,
			http.StatusUnauthorized,
			err.Error(),
			nil,
		)
		return
	}

	token, err := utils.CheckToken(tokenStr)
	if err != nil || !token.Valid {
		utils.WriteErrorResponse(
			c,
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			nil,
		)
		return
	}

	if claims, ok := token.Claims.(*utils.CustomClaim); ok {
		c.Set("user", claims.User)

		c.Next()
	} else {
		utils.WriteErrorResponse(
			c,
			http.StatusUnauthorized,
			http.StatusText(http.StatusUnauthorized),
			nil,
		)
		return
	}
}
