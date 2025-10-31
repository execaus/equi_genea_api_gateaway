package models

type HorseGenderOutput struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetHorseGenderListResponse struct {
	Genders []*HorseGenderOutput `json:"genders"`
}
