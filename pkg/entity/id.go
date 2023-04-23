package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return ID(uuid.New())
}
func NewStrigID() string {
	return NewID().String()
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}
