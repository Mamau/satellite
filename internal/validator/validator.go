package validator

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
)

type Validator struct {
	Translator       ut.Translator
	ValidatorService *validator.Validate
}

func NewValidator() *Validator {
	v := Validator{}
	v.initService()

	return &v
}

// Validate Return slice of errors if exists, and bool is valid
func (v *Validator) Validate(entity interface{}) ([]string, bool) {
	err := v.ValidatorService.Struct(entity)
	v.translateError(err)
	errors := v.translateError(err)
	if len(errors) > 0 {
		return errors, false
	}
	return errors, true
}

func (v *Validator) initService() {
	v.initValidator()
	v.initEnTranslator()
}

func (v *Validator) initValidator() {
	v.ValidatorService = validator.New()
}

func (v *Validator) initEnTranslator() {
	english := en.New()
	uni := ut.New(english, english)
	trans, _ := uni.GetTranslator("en")
	_ = enTranslations.RegisterDefaultTranslations(v.ValidatorService, trans)
	v.Translator = trans
}

func (v *Validator) translateError(err error) []string {
	var errs []string
	if err == nil {
		return errs
	}

	validatorErrs := err.(validator.ValidationErrors)
	for _, e := range validatorErrs {
		translatedErr := e.Translate(v.Translator)
		errs = append(errs, translatedErr)
	}
	return errs
}
