package usecase

import (
	"context"
	"fmt"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type GetAncestorsUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewGetAncestorsUseCase(repository entity.PersonRepositoryInterface) *GetAncestorsUseCase {
	return &GetAncestorsUseCase{
		PersonRepository: repository,
	}
}

func (r *GetAncestorsUseCase) Execute(ctx context.Context, input *dto.GetAncestorsInput) (*dto.GetAncestorsOutput, error) {
	person, err := r.PersonRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return nil, err
	}

	if person == nil {
		return nil, nil
	}

	ancestors, err := r.PersonRepository.FindAncestors(ctx, person)
	if err != nil {
		return nil, err
	}

	output := &dto.GetAncestorsOutput{
		Person: dto.Person{
			Name: person.Name,
			UUID: person.UUID,
		},
		Ancestors: []dto.Ancestors{},
	}

	for _, ancestor := range ancestors {
		parents := ancestor.Relationships.Parents

		fmt.Println("parents:", parents)
		output.Ancestors = append(output.Ancestors, dto.Ancestors{
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
