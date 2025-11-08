package app

import (
	"equi_genea_api_gateaway/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getHorseColorList(c *gin.Context) {
	getColorListResponse, err := h.services.Horse.GetColorList(c, nil)
	if err != nil {
		sendInternalError(c, err)
		return
	}

	colorsOutput := make([]*models.HorseColorOutput, len(getColorListResponse.Colors))
	for i := 0; i < len(colorsOutput); i++ {
		colorsOutput[i] = &models.HorseColorOutput{
			Id:          getColorListResponse.Colors[i].Id,
			Name:        getColorListResponse.Colors[i].Name,
			Description: getColorListResponse.Colors[i].Description,
		}
	}

	sendOK(c, &models.GetHorseColorListResponse{Colors: colorsOutput})
}
