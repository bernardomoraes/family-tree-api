package relationship

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/dto"
	"github.com/bernardomoraes/family-tree/internal/entity"
)

type GetBaconNumberUseCase struct {
	RelationshipRepository entity.RelationshipRepositoryInterface
}

func NewGetBaconNumberUseCase(repository entity.RelationshipRepositoryInterface) *GetBaconNumberUseCase {
	return &GetBaconNumberUseCase{
		RelationshipRepository: repository,
	}
}

func (r *GetBaconNumberUseCase) Execute(ctx context.Context, input *dto.GetBaconNumberInput) (*dto.GetBaconNumberOutput, error) {
	start := entity.Person{
		UUID: input.StartIdentifier,
	}
	end := entity.Person{
		UUID: input.EndIdentifier,
	}

	result, err := r.RelationshipRepository.GetDegreeSeparation(ctx, &start, &end)
	if err != nil {
		return nil, err
	}

	output := &dto.GetBaconNumberOutput{
		BaconNumber: result,
	}

	return output, nil
}
