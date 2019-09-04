package forms

import (
	"net/http"
	"reflect"
)

type Row struct {
	Name     string
	Label    string
	Control  Control
	Tooltip  string
	Required bool
	Readonly bool
	Disabled bool
}

type Error struct {
	Err   error
	Field string
}

type Form struct {
	Errors []Error
	Rows   []Row
	Data   Data
}

func From(v interface{}) (*Form, error) {
	if i, ok := v.(interface{ Form() *Form }); ok {

		return i.Form(), nil
	}
	// TODO: generated interface

	value := reflect.ValueOf(v)
	if value.Kind() != reflect.Ptr || value.Elem().Kind() != reflect.Struct {
		return nil, ErrInvalidDataType
	}

	rb := &reflectBuilder{
		Builder: New(),
		value:   value.Elem(),
	}

	return rb.Build()
}

type reflectBuilder struct {
	*Builder
	value reflect.Value
}

func (rb *reflectBuilder) Build() (*Form, error) {
	t := rb.value.Type()
	for i := 0; i < t.NumField(); i++ {
		if err := rb.addField(t.Field(i)); err != nil {
			return nil, err
		}
	}

	return rb.Builder.Build(NewReflectData(rb.value)), nil
}

func (rb *reflectBuilder) addField(field reflect.StructField) error {
	switch {
	case field.Type.Kind() == reflect.Bool:
		return rb.addBool(field)
	}
}

func (rb *reflectBuilder) addBool(field reflect.StructField) error {
	return nil
}

func (form *Form) Parse(r *http.Request) bool {
	return true
}

func (form *Form) AddFormError(err error) {}

func (form *Form) AddFieldError(name string, err error) {}

func (form *Form) IsValid() bool {
	return false
}
