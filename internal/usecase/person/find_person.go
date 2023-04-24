package usecase

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type FindOnePersonUseCase struct {
	Repository entity.PersonRepositoryInterface
}

func NewFindOnePersonUseCase(repository entity.PersonRepositoryInterface) *FindOnePersonUseCase {
	return &FindOnePersonUseCase{
		Repository: repository,
	}
}

func (f *FindOnePersonUseCase) Execute(ctx context.Context, input *dto.FindPersonInputDTO) (*dto.FindPersonOutputDTO, error) {
	person, err := f.Repository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return nil, err
	}

	if person == nil {
		return nil, nil
	}

	output := &dto.FindPersonOutputDTO{
		Name:          person.Name,
		UUID:          person.UUID,
		Relationships: []dto.Relationship{},
		AuditTrail: dto.AuditTrail{
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
		},
	}

	return output, nil
}