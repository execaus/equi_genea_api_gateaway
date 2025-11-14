package app

import (
	"equi_genea_api_gateaway/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getBreeds(c *gin.Context) {
	getBreedListResponse, err := h.services.Horse.GetBreedList(c, nil)
	if err != nil {
		sendInternalError(c, err)
		return
	}

	breedOutput := make([]*models.HorseBreedOutput, len(getBreedListResponse.Breeds))
	for i := 0; i < len(breedOutput); i++ {
		breedOutput[i] = &models.HorseBreedOutput{
			Id:          getBreedListResponse.Breeds[i].Id,
			Name:        getBreedListResponse.Breeds[i].Name,
			Description: getBreedListResponse.Breeds[i].Description,
		}
	}

	sendOK(c, &models.GetHorseBreedListResponse{Breeds: breedOutput})
}
