package model

import (
	"testing"

	"github.com/brianvoe/gofakeit/v4"
	"github.com/stretchr/testify/assert"
)

func TestValidatorError(t *testing.T) {
	v := Visitor{
		Name: gofakeit.Password(true, true, true, true, true, 32),
	}

	err := v.Validate()

	assert.NotNil(t, err)
}

func TestValidatorPass(t *testing.T) {
	v := Visitor{
		Name: gofakeit.Name(),
	}

	err := v.Validate()

	assert.Nil(t, err)
}
