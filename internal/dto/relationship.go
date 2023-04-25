package dto

type Relationship struct {
	Name         string `json:"name"`
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
}
type CreateParentRelationshipInput struct {
	ParentUUID string `json:"parent" validate:"required"`
	ChildUUID  string `json:"child" validate:"required"`
}

type CreateParentRelationshipByNameInput struct {
	ParentName string `json:"parent_name" validate:"required"`
	ChildName  string `json:"child_name" validate:"required"`
}

type CreateParentRelationshipOutput struct {
	UUID         string `json:"uuid"`
	Relationship string `json:"relationship"`
	RelatedUUID  string `json:"related_uuid"`
}

type CreateParentRelationshipBulkInput struct {
	Relationships []CreateParentRelationshipInput `json:"relationships" validate:"required"`
}
