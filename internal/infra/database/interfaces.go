package database

import (
	"context"

	"github.com/bernardomoraes/family-tree/internal/entity"
)

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}

type PersonInterface interface {
	Create(ctx context.Context, person *entity.Person) (*entity.Person, error)
	FindByUUID(ctx context.Context, uuid string) (*entity.Person, error)
	Update(ctx context.Context, person *entity.Person) (*entity.Person, error)
	// FindByEmail(ctx context.Context, email string) (*entity.Person, error)
}
