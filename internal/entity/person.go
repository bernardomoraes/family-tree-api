package entity

import (
	"errors"
	"time"

	"github.com/bernardomoraes/family-tree/pkg/entity"
)

var (
	ErrIDIsRequired   = errors.New("id is required")
	ErrInvalidID      = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
)

type Person struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewPerson(name string) (*Person, error) {
	product := &Person{
		ID:        entity.NewID(),
		Name:      name,
		CreatedAt: time.Now(),
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Person) Validate() error {
	if p.ID.String() == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}
	return nil
}
