package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPerson(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Equal(t, p.Name, "John Doe")
	assert.NotEmpty(t, p.UUID)
	assert.Equal(t, "John Doe", p.Name)
}

func TestPerson_NameIsRequired(t *testing.T) {
	p, err := NewPerson("")

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestPerson_UUIDIsRequired(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.UUID)
}

func TestPerson_Validate(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}

func TestPerson_Validate_BlankUUID(t *testing.T) {
	p, err := NewPerson("John Doe")
	p.UUID = ""

	assert.Nil(t, err)
	assert.ErrorIs(t, p.Validate(), ErrIDIsRequired)
}

func TestPerson_Validate_InvalidUUID(t *testing.T) {
	p, err := NewPerson("John Doe")
	p.UUID = "123"

	assert.Nil(t, err)
	assert.ErrorIs(t, p.Validate(), ErrInvalidID)
}
