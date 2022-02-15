package controllers

import "github.com/gin-gonic/gin"

type ErrorResponse struct {
	Error string `json:"error"`
	Code  int    `json:"code"`
}

func errorHandler(c *gin.Context, message string, code int) {
	c.JSON(code, ErrorResponse{
		Error: message,
		Code:  code,
	})
}
