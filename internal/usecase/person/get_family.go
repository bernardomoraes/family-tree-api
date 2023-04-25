package usecase

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type GetFamilyUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewGetFamilyUseCase(repository entity.PersonRepositoryInterface) *GetFamilyUseCase {
	return &GetFamilyUseCase{
		PersonRepository: repository,
	}
}

func (r *GetFamilyUseCase) Execute(ctx context.Context, input *dto.GetAncestorsInput) (*dto.GetFamilyOutput, error) {
	person, err := r.PersonRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return nil, err
	}

	if person == nil {
		return nil, nil
	}

	Family, err := r.PersonRepository.FindFamily(ctx, person)
	if err != nil {
		return nil, err
	}

	output := &dto.GetFamilyOutput{
		Person: dto.Person{
			Name: person.Name,
			UUID: person.UUID,
		},
		Family: []dto.Ancestors{},
	}

	for _, ancestor := range Family {
		output.Family = append(output.Family, dto.Ancestors{
			Person: dto.Person{
				Name: ancestor.Name,
				UUID: ancestor.UUID,
			},
			Relationships: &dto.RelationshipList{
				Parents: ancestor.Relationships.Parents,
				Childs:  ancestor.Relationships.Childs,
			},
		})
	}

	return output, nil
}
