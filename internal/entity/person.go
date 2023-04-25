package entity

import (
	"errors"

	"github.com/bernardomoraes/family-tree/pkg/entity"
)

var (
	ErrIDIsRequired   = errors.New("id is required")
	ErrInvalidID      = errors.New("invalid id")
	ErrNameIsRequired = errors.New("name is required")
)

type Relationships struct {
	Childs  []Person `json:"childs,omitempty"`
	Parents []Person `json:"parents,omitempty"`
}
type AuditTrail struct {
	CreatedAt string `json:"createdAt,omitempty"`
	UpdatedAt string `json:"updatedAt,omitempty"`
}

type Person struct {
	ID             int64  `json:"id,omitempty"`
	UUID           string `json:"uuid"`
	Name           string `json:"name"`
	*Relationships `json:"relationships,omitempty"`
	AuditTrail
}

func NewPerson(name string) (*Person, error) {
	person := &Person{
		UUID: entity.NewStrigID(),
		Name: name,
	}

	err := person.Validate()
	if err != nil {
		return nil, err
	}

	return person, nil
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
