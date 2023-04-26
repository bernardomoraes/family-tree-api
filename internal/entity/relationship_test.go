package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewRelationship(t *testing.T) {
	r, err := NewRelationship("John Doe", "Jane Doe", "IS_PARENT")

	assert.Nil(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, r.Start, "John Doe")
	assert.Equal(t, r.End, "Jane Doe")
	assert.Equal(t, r.Relation, "IS_PARENT")
	assert.NotEmpty(t, r.Relation)
	assert.NotEmpty(t, r.End)
}

func TestRelationship_NameIsRequired(t *testing.T) {
	p, err := NewRelationship("", "", "IS_PARENT")

	assert.Nil(t, p)
	assert.Equal(t, ErrStartAndEndIsRequired, err)
}

func TestRelationship_RelationTypeIsRequired(t *testing.T) {
	p, err := NewRelationship("John Doe", "Jane Doe", "")

	assert.EqualError(t, err, ErrRelTypeIsRequired.Error())
	assert.Nil(t, p)
}

func TestRelationship_RelationTypeIsWrong(t *testing.T) {
	p, err := NewRelationship("John Doe", "Jane Doe", "IS_NONE")

	assert.EqualError(t, err, ErrInvalidType.Error())
	assert.Nil(t, p)
}

// func TestPersonValidate(t *testing.T) {
// 	p, err := NewRelationship("John Doe")

// 	assert.Nil(t, err)
// 	assert.NotNil(t, p)
// 	assert.Nil(t, p.Validate())
// }
