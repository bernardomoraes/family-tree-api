package relationship

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type CreateRelationshipUseCase struct {
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewCreateRelationshipUseCase(repository entity.RelationshipRepositoryInterface) *CreateRelationshipUseCase {
	return &CreateRelationshipUseCase{
		RelationshipRepository: repository,
	}
}

func (r *CreateRelationshipUseCase) Execute(ctx context.Context, input *dto.CreateParentRelationshipInput) error {
	parent := entity.Person{
		UUID: input.ParentUUID,
	}
	child := entity.Person{
		UUID: input.ChildUUID,
	}
	relationship, err := entity.NewRelationship(parent.UUID, child.UUID, "IS_PARENT")

	if err != nil {
		return err
	}

	rel, err := r.RelationshipRepository.FindRelationship(ctx, relationship)
	if err != nil {
		return err
	}

	if rel != nil {
		return entity.ErrRelationshipAlreadyExists
	}

	err = r.RelationshipRepository.CreateIsParent(ctx, parent, child)
	if err != nil {
		return err
	}

	return nil
}
