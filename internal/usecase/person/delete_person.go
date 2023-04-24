package usecase

import (
	"context"
	"errors"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type DeletePersonUseCase struct {
	PersonRepository entity.PersonRepositoryInterface
}

func NewDeletePersonUseCase(personRepository entity.PersonRepositoryInterface) *DeletePersonUseCase {
	return &DeletePersonUseCase{
		PersonRepository: personRepository,
	}
}

func (c *DeletePersonUseCase) Execute(ctx context.Context, input *dto.DeletePersonInputDTO) error {
	person, err := c.PersonRepository.FindByUUID(ctx, input.UUID)
	if err != nil {
		return err
	}
	if person == nil {
		return errors.New("person not found")
	}

	err = c.PersonRepository.Delete(ctx, person.UUID)
	if err != nil {
		return err
	}

	return nil
}
