package models

type HorseColorOutput struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetHorseColorListResponse struct {
	Colors []*HorseColorOutput `json:"colors"`
}
