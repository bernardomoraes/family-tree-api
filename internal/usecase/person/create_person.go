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

	// Aqui a gente insere as regras de negócio
	// Exemplo: Verificar se tem filho ou pai, verificar se é esposa de alguém e por ai vai

	// Parte de salvar
	person, err = c.PersonRepository.Create(ctx, person)
	if err != nil {
		return nil, err
	}

	output := &dto.CreatePersonOutputDTO{
		Person: dto.Person{
			Name: person.Name,
			UUID: person.UUID,
		},
		AuditTrail: dto.AuditTrail{
			CreatedAt: person.CreatedAt,
			UpdatedAt: person.UpdatedAt,
		},
	}

	return output, nil
}
