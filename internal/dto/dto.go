package dto

type AuditTrail struct {
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type Relationship struct {
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
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
