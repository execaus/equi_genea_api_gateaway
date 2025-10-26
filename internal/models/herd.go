package models

type HerdOutput struct {
	// TODO
}

type CreateHerdRequest struct {
	Name        string  `json:"name" validate:"required,min=4,max=64"`
	Description *string `json:"description" validate:"omitempty,max=1024"`
}

type CreateHerdResponse struct {
	Herd HerdOutput `json:"herd"`
}
