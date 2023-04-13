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
	assert.NotEmpty(t, p.ID)
	assert.NotEmpty(t, p.CreatedAt)
	assert.Equal(t, "John Doe", p.Name)
}

func TestPersonWhenNameIsRequired(t *testing.T) {
	p, err := NewPerson("")

	assert.Nil(t, p)
	assert.Equal(t, ErrNameIsRequired, err)
}

func TestPersonWhenIDIsRequired(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p.ID)
}

func TestPersonValidate(t *testing.T) {
	p, err := NewPerson("John Doe")

	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.Nil(t, p.Validate())
}
