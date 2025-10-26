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
			Password:     generatePasswordResponse.Password,
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

func (h *Handler) signIn(c *gin.Context) {
	var input models.SignInRequest

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

	if !isExistResponse.IsExist {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	getAccountByEmailResponse, err := h.services.Account.GetAccountByEmail(
		c.Request.Context(),
		&account.GetAccountByEmailRequest{
			Email: input.Email,
		},
	)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	comparePasswordResponse, err := h.services.Auth.ComparePassword(c.Request.Context(), &authpb.ComparePasswordRequest{
		HashedPassword: getAccountByEmailResponse.Account.Password,
		Password:       input.Password,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	if !comparePasswordResponse.IsMatch {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	generateTokenResponse, err := h.services.Auth.GenerateToken(c.Request.Context(), &authpb.GenerateTokenRequest{
		Id: getAccountByEmailResponse.Account.Id,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &models.SignInResponse{
		Token: generateTokenResponse.Token,
	})
}
