package dto

type FindRelationshipInput struct {
	UUID         string `json:"uuid" validate:"required"`
	Relationship string `json:"relationship" validate:"required"`
}

type FindRelationshipOutput struct {
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
	RelatedUUID  string `json:"related_uuid"`
}

type ErrorBody struct {
	Message string `json:"message"`
}
