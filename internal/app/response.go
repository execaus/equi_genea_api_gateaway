package app

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorMessage struct {
	Message string `json:"message"`
}

func sendBadRequest(c *gin.Context, err error) {
	c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
}

func sendInternalError(c *gin.Context, err error) {
	fmt.Println(err.Error())
	c.AbortWithStatus(http.StatusInternalServerError)
}

func sendConflict(c *gin.Context, message string) {
	c.AbortWithStatusJSON(http.StatusConflict, &ErrorMessage{Message: message})
}

func sendCreated(c *gin.Context, body any) {
	c.AbortWithStatusJSON(http.StatusCreated, body)
}

func sendOK(c *gin.Context, body any) {
	c.AbortWithStatusJSON(http.StatusOK, body)
}

func sendUnauthorized(c *gin.Context) {
	c.AbortWithStatus(http.StatusUnauthorized)
}
