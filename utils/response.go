package utils

import "github.com/gin-gonic/gin"

type JsonResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func WriteResponse(c *gin.Context, code int, message string, data interface{}) {
	c.JSON(code, JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func WriteErrorResponse(c *gin.Context, code int, message string, data interface{}) {
	c.AbortWithStatusJSON(code, JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
