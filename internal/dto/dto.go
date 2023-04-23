package dto

import "time"

type CreatePersonInput struct {
	Name string `json:"name" validate:"required"`
}

type CreatePersonOutput struct {
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	CreatedAt string `json:"created_at"`
}

type FindPersonInput struct {
	UUID string `json:"uuid" validate:"required"`
}

type Relationship struct {
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
}

type FindPersonOutput struct {
	Name          string         `json:"name"`
	UUID          string         `json:"uuid"`
	Relationships []Relationship `json:"relationships"`
}

type UpdatePersonInput struct {
	Name string `json:"name" validate:"required"`
}

type UpdatePersonOutput struct {
	Name      string    `json:"name"`
	UUID      string    `json:"uuid"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
}

type CreateRelationshipInput struct {
	UUID          string `json:"uuid" validate:"required"`
	Relationship  string `json:"relationship" validate:"required"`
	RelatedPerson string `json:"related_person" validate:"required"`
}

type CreateRelationshipOutput struct {
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
	RelatedUUID  string `json:"related_uuid"`
}

type FindRelationshipInput struct {
	UUID         string `json:"uuid" validate:"required"`
	Relationship string `json:"relationship" validate:"required"`
}

type FindRelationshipOutput struct {
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
	RelatedUUID  string `json:"related_uuid"`
}
