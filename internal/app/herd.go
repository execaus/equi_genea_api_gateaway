package app

import (
	"equi_genea_api_gateaway/internal/models"
	"equi_genea_api_gateaway/internal/pb/api/herd"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createHerd(c *gin.Context) {
	var input models.CreateHerdRequest

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	createHerdResponse, err := h.services.Herd.CreateHerd(c.Request.Context(), &herd.CreateHerdRequest{
		Name:        input.Name,
		Description: input.Description,
		AccountId:   accountID,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	herdOutput := models.HerdOutput{}

	herdOutput.LoadFromHerdPB(createHerdResponse.Herd)

	c.JSON(http.StatusOK, &models.CreateHerdResponse{Herd: herdOutput})
}
