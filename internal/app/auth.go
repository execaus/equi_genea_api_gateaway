package app

import (
	"equi_genea_api_gateaway/internal/models"
	"equi_genea_api_gateaway/internal/pb/api/account"
	authpb "equi_genea_api_gateaway/internal/pb/api/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) signUp(c *gin.Context) {
	var input models.SignUpRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	isExistResponse, err := h.services.Account.IsExistByEmail(
		c.Request.Context(),
		&account.IsExistByEmailRequest{Email: input.Email},
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if isExistResponse.IsExist {
		c.AbortWithStatus(http.StatusConflict)
		return
	}

	generatePasswordResponse, err := h.services.Auth.GeneratePassword(c.Request.Context(), nil)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	hashPasswordResponse, err := h.services.Auth.HashPassword(c.Request.Context(), &authpb.HashPasswordRequest{
		Password: generatePasswordResponse.Password,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	createAccountResponse, err := h.services.Account.CreateAccount(
		c.Request.Context(),
		&account.CreateAccountRequest{
			Email:        input.Email,
			PasswordHash: hashPasswordResponse.Hash,
		},
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	generateTokenResponse, err := h.services.Auth.GenerateToken(
		c.Request.Context(),
		&authpb.GenerateTokenRequest{Id: createAccountResponse.Account.Id},
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &models.SignUpResponse{
		Token: generateTokenResponse.Token,
	})
}
