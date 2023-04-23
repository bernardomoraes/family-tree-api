package usecase

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/entity"
)

type CreatePersonInputDTO struct {
	Name string `json:"name" validate:"required"`
}

type CreatePersonOutputDTO struct {
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	CreatedAt string `json:"created_at"`
}

type CreatePersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewCreatePersonUseCase(personRepository entity.PersonRepositoryInterface) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		PersonRepository: personRepository,
	}
}

func (c *CreatePersonUseCase) Execute(ctx context.Context, input *CreatePersonInputDTO) (*CreatePersonOutputDTO, error) {
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

	output := &CreatePersonOutputDTO{
		Name:      person.Name,
		UUID:      person.UUID,
		CreatedAt: person.CreatedAt.String(),
	}

	return output, nil
}
