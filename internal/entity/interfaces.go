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
	FindByName(ctx context.Context, name string) (*Person, error)
	Update(ctx context.Context, person *Person) (*Person, error)
	Delete(ctx context.Context, uuid string) error
	FindAncestors(ctx context.Context, person *Person) ([]*Person, error)
}

type RelationshipRepositoryInterface interface {
	CreateIsParent(ctx context.Context, parent Person, child Person) error
	FindRelationship(ctx context.Context, relationship *Relationship) (*Relationship, error)
	GetDegreeSeparation(ctx context.Context, person1 *Person, person2 *Person) (int, error)
}
