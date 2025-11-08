package app

import (
	"equi_genea_api_gateaway/internal/models"
	herdpb "equi_genea_api_gateaway/internal/pb/api/herd"
	"errors"

	"github.com/gin-gonic/gin"
)

func (h *Handler) createHerd(c *gin.Context) {
	var input models.CreateHerdRequest

	if err := c.BindJSON(&input); err != nil {
		sendBadRequest(c, err)
		return
	}

	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		sendUnauthorized(c)
		return
	}

	createHerdResponse, err := h.services.Herd.CreateHerd(c.Request.Context(), &herdpb.CreateHerdRequest{
		Name:        input.Name,
		Description: input.Description,
		AccountId:   accountID,
	})
	if err != nil {
		sendInternalError(c, err)
		return
	}

	herdOutput := models.HerdOutput{}

	herdOutput.LoadFromHerdPB(createHerdResponse.Herd)

	sendOK(c, &models.CreateHerdResponse{Herd: herdOutput})
}

func (h *Handler) getHerdList(c *gin.Context) {
	var queryParams models.GetListParams

	queryParams.BindFromContext(c)

	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		sendUnauthorized(c)
		return
	}

	getHerdListResponse, err := h.services.Herd.GetHerdList(c.Request.Context(), &herdpb.GetHerdListRequest{
		Limit:     queryParams.Limit,
		Page:      queryParams.Page,
		Search:    queryParams.Search,
		AccountId: accountID,
	})
	if err != nil {
		sendInternalError(c, err)
		return
	}

	herds := make([]models.HerdOutput, len(getHerdListResponse.Herds))
	for i, herd := range getHerdListResponse.Herds {
		herds[i] = models.HerdOutput{}
		herds[i].LoadFromHerdPB(herd)
	}

	sendOK(c, &models.GetHerdListResponse{
		Herds:      herds,
		TotalCount: getHerdListResponse.TotalCount,
	})
}

func (h *Handler) getHerdByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		sendBadRequest(c, errors.New("invalid path parameter id"))
		return
	}

	getHerdByIdResponse, err := h.services.Herd.GetHerdById(c.Request.Context(), &herdpb.GetHerdByIdRequest{Id: id})
	if err != nil {
		sendInternalError(c, err)
		return
	}

	herd := models.HerdOutput{}
	herd.LoadFromHerdPB(getHerdByIdResponse.Herd)

	sendOK(c, &models.GetHerdByIDResponse{Herd: &herd})
}
