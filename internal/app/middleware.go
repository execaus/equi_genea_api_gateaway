package app

import (
	"equi_genea_api_gateaway/internal/pb/api/auth"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

type contextKey string

const AccountIDKey contextKey = "accountID"

func (h *Handler) authMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		fmt.Println("missing or invalid Authorization header")
		sendUnauthorized(c)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")

	getClaimsFromTokenResponse, err := h.services.Auth.GetClaimsFromToken(c.Request.Context(), &auth.GetClaimsFromTokenRequest{
		Token: token,
	})
	if err != nil {
		fmt.Println(err.Error())
		sendUnauthorized(c)
		return
	}

	c.Set(AccountIDKey, getClaimsFromTokenResponse.Claims.AccountId)
	c.Next()
}

func getAccountIDFromContext(ctx *gin.Context) (string, bool) {
	accountID, ok := ctx.Get(AccountIDKey)
	if !ok {
		return "", false
	}
	idStr, ok := accountID.(string)
	return idStr, ok
}
