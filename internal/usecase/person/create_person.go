package usecase

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type CreatePersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewCreatePersonUseCase(personRepository entity.PersonRepositoryInterface) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		PersonRepository: personRepository,
	}
}

func (c *CreatePersonUseCase) Execute(ctx context.Context, input *dto.CreatePersonInputDTO) (*dto.CreatePersonOutputDTO, error) {
	person, err := entity.NewPerson(input.Name)
	if err != nil {
		return nil, err
	}

	// Parte de salvar
	personCreated, err := c.PersonRepository.Create(ctx, person)
	if err != nil {
		return nil, err
	}

	output := &dto.CreatePersonOutputDTO{
		Person: dto.Person{
			Name: personCreated.Name,
			UUID: personCreated.UUID,
		},
		AuditTrail: dto.AuditTrail{
			CreatedAt: personCreated.CreatedAt,
			UpdatedAt: personCreated.UpdatedAt,
		},
	}

	return output, nil
}
