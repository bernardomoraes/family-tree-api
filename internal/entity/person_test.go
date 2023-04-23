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
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, "John Doe", p.Name)
}

func TestPersonWhenNameIsRequired(t *testing.T) {
	p, err := NewPerson("")

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestPersonWhenUUIDIsRequired(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.UUID)
}

func TestPersonValidate(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
