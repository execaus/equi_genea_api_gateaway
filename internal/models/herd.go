package models

import (
	"equi_genea_api_gateaway/internal/pb/api/herd"
	"time"
)

type HerdOutput struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
	AccountID   string    `json:"account_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (h *HerdOutput) LoadFromHerdPB(pb *herd.Herd) {
	if pb == nil {
		return
	}

	h.ID = pb.Id
	h.Name = pb.Name
	h.Description = pb.Description
	h.AccountID = pb.AccountId

	if pb.CreatedAt != nil {
		h.CreatedAt = pb.CreatedAt.AsTime()
	}
	if pb.UpdatedAt != nil {
		h.UpdatedAt = pb.UpdatedAt.AsTime()
	}
}

type CreateHerdRequest struct {
	Name        string  `json:"name" validate:"required,min=4,max=64"`
	Description *string `json:"description" validate:"omitempty,max=1024"`
}

type CreateHerdResponse struct {
	Herd HerdOutput `json:"herd"`
}

type GetHerdListResponse struct {
	Herds      []HerdOutput `json:"herds"`
	TotalCount int32        `json:"totalCount"`
}

type GetHerdByIDResponse struct {
	Herd *HerdOutput `json:"herd"`
}
