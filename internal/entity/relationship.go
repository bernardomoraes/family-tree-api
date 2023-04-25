package entity

import (
	"errors"
)

var (
	ErrRelTypeIsRequired         = errors.New("Relationship type is required")
	ErrInvalidType               = errors.New("Relationship type is invalid")
	ErrStartAndEndIsRequired     = errors.New("Relationship start and end are required")
	ErrRelationshipAlreadyExists = errors.New("Relationship already exists")
)

type Relationship struct {
	ID       int64  `json:"id"`
	Relation string `json:"relation" validate:"required"`
	Start    string `json:"start" validate:"required"`
	End      string `json:"end" validate:"required"`
}

func NewRelationship(start string, end string, relType string) (*Relationship, error) {
	relationship := &Relationship{
		Relation: relType,
		Start:    start,
		End:      end,
	}

	err := relationship.Validate()
	if err != nil {
		return nil, err
	}

	return relationship, nil
}

func (r *Relationship) Validate() error {
	if r.Relation == "" {
		return ErrRelTypeIsRequired
	}

	categories := []string{"IS_PARENT", "IS_SPAUSE"}
	if !contains(categories, r.Relation) {
		return ErrInvalidType
	}

	if r.Start == "" || r.End == "" {
		return ErrStartAndEndIsRequired
	}

	return nil
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
