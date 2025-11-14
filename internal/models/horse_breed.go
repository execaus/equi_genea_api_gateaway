package models

type HorseBreedOutput struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetHorseBreedListResponse struct {
	Breeds []*HorseBreedOutput `json:"breeds"`
}
