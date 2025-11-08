package models

type HorseGeneticMarkerOutput struct {
	Id          string  `json:"id"`
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type GetHorseGeneticMarkerListResponse struct {
	GeneticMarkers []*HorseGeneticMarkerOutput `json:"geneticMarkers"`
}
