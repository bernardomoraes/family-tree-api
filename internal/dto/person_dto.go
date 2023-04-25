package dto

import "github.com/bernardomoraes/family-tree/internal/entity"

type Person entity.Person
type AuditTrail entity.AuditTrail
type RelationshipList entity.Relationships
type CreatePersonInputDTO struct {
	Name string `json:"name" validate:"required"`
}

type CreatePersonOutputDTO struct {
	Person
	AuditTrail
}

type FindPersonInputDTO struct {
	Person
}

type FindPersonOutputDTO struct {
	Person
	Relationships []Relationship `json:"relationships"`
	AuditTrail
}

type UpdatePersonInputDTO struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePersonOutputDTO struct {
	Person
	AuditTrail
}

type DeletePersonInputDTO struct {
	UUID string `json:"uuid" validate:"required"`
}

type GetAncestorsInput struct {
	UUID string `json:"uuid" validate:"required"`
}

type Ancestors struct {
	Person
	Relationships *RelationshipList `json:"relationships,omitempty"`
}
type GetAncestorsOutput struct {
	Person
	Relationships *RelationshipList `json:"relationships,omitempty"`
	Ancestors     []Ancestors       `json:"ancestors"`
}
