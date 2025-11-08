package models

type HorseBirthplaceOutput struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetHorseBirthplaceListResponse struct {
	Birthplaces []*HorseBirthplaceOutput `json:"birthplaces"`
}
