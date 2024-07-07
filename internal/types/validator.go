package types

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var (
	Validate    = validator.New()
	enLocales   = en.New()
	Translators = ut.New(enLocales)
)

func init() {
	trans, _ := Translators.GetTranslator("en")

	en_translations.RegisterDefaultTranslations(Validate, trans)

	Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "{0} is required", true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("required", fe.Field())
		return t
	})
}
