package app

import (
	"equi_genea_api_gateaway/internal/models"
	herdpb "equi_genea_api_gateaway/internal/pb/api/herd"
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

	createHerdResponse, err := h.services.Herd.CreateHerd(c.Request.Context(), &herdpb.CreateHerdRequest{
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

func (h *Handler) getHerdList(c *gin.Context) {
	var queryParams models.GetListParams

	queryParams.BindFromContext(c)

	accountID, exists := getAccountIDFromContext(c)
	if !exists {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	getHerdListResponse, err := h.services.Herd.GetHerdList(c.Request.Context(), &herdpb.GetHerdListRequest{
		Limit:     queryParams.Limit,
		Page:      queryParams.Page,
		Search:    queryParams.Search,
		AccountId: accountID,
	})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	herds := make([]models.HerdOutput, len(getHerdListResponse.Herds))
	for i, herd := range getHerdListResponse.Herds {
		herds[i] = models.HerdOutput{}
		herds[i].LoadFromHerdPB(herd)
	}

	c.JSON(http.StatusOK, &models.GetHerdListResponse{
		Herds:      herds,
		TotalCount: getHerdListResponse.TotalCount,
	})
}

func (h *Handler) getHerdByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	getHerdByIdResponse, err := h.services.Herd.GetHerdById(c.Request.Context(), &herdpb.GetHerdByIdRequest{Id: id})
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	herd := models.HerdOutput{}
	herd.LoadFromHerdPB(getHerdByIdResponse.Herd)

	c.JSON(http.StatusOK, &models.GetHerdByIDResponse{Herd: &herd})
}
