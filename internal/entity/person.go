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
	ID        int64     `json:"id"`
	UUID      string    `json:"uuid"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewPerson(name string) (*Person, error) {
	product := &Person{
		UUID: entity.NewStrigID(),
		Name: name,
	}

	err := product.Validate()
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Person) Validate() error {
	if p.UUID == "" {
		return ErrIDIsRequired
	}

	if _, err := entity.ParseID(p.UUID); err != nil {
		return ErrInvalidID
	}

	if p.Name == "" {
		return ErrNameIsRequired
	}
	return nil
}
