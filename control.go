package forms

// Control an abstract form control
type Control interface {
	Kind() string
	Parse(string) (interface{}, error)
	GetValue(interface{}) string
}

// CheckboxControl is an boolean form value
type CheckboxControl struct {
	Value string
}

// Checkbox is an boolean form control
func Checkbox(v string) CheckboxControl {
	if v == "" {
		v = "true"
	}

	return CheckboxControl{v}
}

// Kind implements Control interface
func (c CheckboxControl) Kind() string {
	return "checkbox"
}

// Parse implements Control interface
func (c CheckboxControl) Parse(s string) (interface{}, error) {
	return s == c.Value, nil
}

// Value implements Control interface
func (c CheckboxControl) GetValue(v interface{}) string {
	if b, ok := v.(bool); ok && b {
		return c.Value
	}
	return ""
}

// type Number struct{}

// type Range struct{}

// type String struct{}

// type Password struct{}

// type Text struct{}

// type Password struct{}

// type Email struct{}

// type Tel struct{}

// type Hidden struct{}

// type Select struct{}

// type Radio struct{}

// type Datetime struct{}

// type File struct{}
