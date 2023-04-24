package dto

type CreatePersonInputDTO struct {
	Name string `json:"name" validate:"required"`
}

type CreatePersonOutputDTO struct {
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	CreatedAt string `json:"created_at"`
}
