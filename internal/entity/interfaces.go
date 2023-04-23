package entity

import (
	"context"
)

type UserRepositoryInterface interface {
	Create(user *User) error
	FindByEmail(email string) (*User, error)
}

type PersonRepositoryInterface interface {
	Create(ctx context.Context, person *Person) (*Person, error)
	FindByUUID(ctx context.Context, uuid string) (*Person, error)
	Update(ctx context.Context, person *Person) (*Person, error)
	Delete(ctx context.Context, uuid string) error
}

type RelationshipRepositoryInterface interface {
	Create(ctx context.Context, relationship *Relationship) (*Relationship, error)
	Update(ctx context.Context, relationship *Relationship) (*Relationship, error)
	Delete(ctx context.Context, uuid string) error
}
