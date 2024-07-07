package forms

import (
	"lisfun/internal/types"
	"strings"

	"github.com/go-playground/validator/v10"
)

type Form struct {
	Values map[string]any
	Errors map[string]string
}

func (f *Form) Value(field string) string {
	if f == nil {
		return ""
	}

	value, ok := f.Values[field]
	if !ok {
		return ""
	}

	stringValue, _ := value.(string)
	return stringValue
}

func (f *Form) HasError(field string) bool {
	if f == nil {
		return false
	}

	_, ok := f.Errors[field]
	return ok
}

func (f *Form) Error(field string) string {
	if f == nil {
		return ""
	}

	errorString := f.Errors[field]
	return errorString
}

func removeTopStruct(fields map[string]string) map[string]string {
	res := map[string]string{}
	for field, err := range fields {
		res[field[strings.Index(field, ".")+1:]] = err
	}
	return res
}

func FromValidationsErrors(err error, lang string) (*Form, bool) {
	ves, ok := err.(validator.ValidationErrors)
	if !ok {
		return nil, false
	}

	form := &Form{
		Values: map[string]any{},
		Errors: map[string]string{},
	}

	trans, _ := types.Translators.GetTranslator(lang)
	form.Errors = removeTopStruct(ves.Translate(trans))

	for _, ve := range ves {
		form.Values[ve.Field()] = ve.Value()
	}

	return form, true
}
