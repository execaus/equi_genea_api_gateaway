package app

import (
	"equi_genea_api_gateaway/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getGeneticMarkers(c *gin.Context) {
	getGeneticMarkerListResponse, err := h.services.Horse.GetGeneticMarkerList(c, nil)
	if err != nil {
		sendInternalError(c, err)
		return
	}

	markersOutput := make([]*models.HorseGeneticMarkerOutput, len(getGeneticMarkerListResponse.Markers))
	for i := 0; i < len(markersOutput); i++ {
		markersOutput[i] = &models.HorseGeneticMarkerOutput{
			Id:          getGeneticMarkerListResponse.Markers[i].Id,
			Name:        getGeneticMarkerListResponse.Markers[i].Name,
			Description: getGeneticMarkerListResponse.Markers[i].Description,
		}
	}

	sendOK(c, &models.GetHorseGeneticMarkerListResponse{GeneticMarkers: markersOutput})
}
