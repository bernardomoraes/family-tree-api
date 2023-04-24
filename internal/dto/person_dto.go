package dto

type CreatePersonInputDTO struct {
	Name string `json:"name" validate:"required"`
}

type CreatePersonOutputDTO struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	AuditTrail
}

type FindPersonInputDTO struct {
	UUID string `json:"uuid" validate:"required"`
}

type FindPersonOutputDTO struct {
	Name          string         `json:"name"`
	UUID          string         `json:"uuid"`
	Relationships []Relationship `json:"relationships"`
	AuditTrail
}

type UpdatePersonInputDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePersonOutputDTO struct {
	Name string `json:"name"`
	UUID string `json:"uuid"`
	AuditTrail
}
