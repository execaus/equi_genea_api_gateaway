package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createHerd(c *gin.Context) {
	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_ = accountID
}
