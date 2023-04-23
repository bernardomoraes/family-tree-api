package entity

import (
	"errors"

	"github.com/bernardomoraes/family-tree/pkg/helpers"
)

var (
	ErrRelTypeIsRequired = errors.New("Relationship type is required")
	ErrInvalidType       = errors.New("Relationship type is invalid")
)

type Relationship struct {
	ID   int64  `json:"id"`
	Type string `json:"type"`
}

func NewReleationship(relType string) (*Relationship, error) {
	relationship := &Relationship{
		Type: relType,
	}

	err := relationship.Validate()
	if err != nil {
		return nil, err
	}

	return relationship, nil
}

func (p *Relationship) Validate() error {
	if p.Type == "" {
		return ErrRelTypeIsRequired
	}

	categories := []string{"parent", "spouse"}
	if !helpers.Contains(categories, p.Type) {
		return ErrInvalidType
	}

	return nil
}
