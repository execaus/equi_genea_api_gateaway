package app

import (
	"equi_genea_api_gateaway/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getHorseGenderList(c *gin.Context) {
	getGenderListResponse, err := h.services.Horse.GetGenderList(c, nil)
	if err != nil {
		sendInternalError(c, err)
		return
	}

	gendersOutput := make([]*models.HorseGenderOutput, len(getGenderListResponse.Genders))
	for i := 0; i < len(gendersOutput); i++ {
		gendersOutput[i] = &models.HorseGenderOutput{
			Id:          getGenderListResponse.Genders[i].Id,
			Name:        getGenderListResponse.Genders[i].Name,
			Description: getGenderListResponse.Genders[i].Description,
		}
	}

	sendOK(c, &models.GetHorseGenderListResponse{Genders: gendersOutput})
}
