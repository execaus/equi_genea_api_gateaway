package models

type SignUpRequest struct {
	Email string `json:"email" validate:"required,email,min=4,max=255"`
}

type SignUpResponse struct {
	Token string `json:"token"`
}
