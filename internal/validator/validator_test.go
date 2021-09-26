package validator

import (
	"testing"

	"satellite/internal/entity"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

func TestValidate(t *testing.T) {
	v := NewValidator()
	r := entity.Run{}
	errorsList, isValid := v.Validate(r)
	assert.False(t, isValid)
	assert.NotEmpty(t, errorsList)
	assert.Equal(t, "Name is a required field", errorsList[0])
	assert.Equal(t, "Image is a required field", errorsList[1])

	r.Name = "test"
	r.Image = "image"
	errorsList, isValid = v.Validate(r)
	assert.True(t, isValid)
	assert.Empty(t, errorsList)
}

func TestTranslateError(t *testing.T) {
	v := NewValidator()
	r := entity.Run{}
	err := v.ValidatorService.Struct(r)

	result := v.translateError(err)
	assert.NotEmpty(t, result)
	assert.Equal(t, "Name is a required field", result[0])
	assert.Equal(t, "Image is a required field", result[1])

	r.Name = "test"
	r.Image = "image"
	err = v.ValidatorService.Struct(r)
	result = v.translateError(err)
	assert.Empty(t, result)
}

func TestInitEnTranslator(t *testing.T) {
	v := NewValidator()
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")

	assert.IsType(t, trans, v.Translator)
	assert.ObjectsAreEqual(trans, v.Translator)
}

func TestInitValidator(t *testing.T) {
	v := NewValidator()
	assert.IsType(t, validator.New(), v.ValidatorService)
	assert.ObjectsAreEqual(validator.New(), v.ValidatorService)
}
