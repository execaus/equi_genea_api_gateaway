package app

import (
	"equi_genea_api_gateaway/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getHorseBirthplaces(c *gin.Context) {
	getBirthplaceListResponse, err := h.services.Horse.GetBirthplaceList(c, nil)
	if err != nil {
		sendInternalError(c, err)
		return
	}

	birthplaceOutput := make([]*models.HorseBirthplaceOutput, len(getBirthplaceListResponse.Birthplaces))
	for i := 0; i < len(birthplaceOutput); i++ {
		birthplaceOutput[i] = &models.HorseBirthplaceOutput{
			Id:          getBirthplaceListResponse.Birthplaces[i].Id,
			Name:        getBirthplaceListResponse.Birthplaces[i].Name,
			Description: getBirthplaceListResponse.Birthplaces[i].Description,
		}
	}

	sendOK(c, &models.GetHorseBirthplaceListResponse{Birthplaces: birthplaceOutput})
}
