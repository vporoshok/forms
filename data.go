package forms

import "reflect"

// Data is an interface to store values from form
type Data interface {
	Set(name string, value interface{})
}

// MapData is an wrapper of hash-table usefull for tests
type MapData map[string]interface{}

// Set implements Data interface
func (md *MapData) Set(name string, value interface{}) {
	(*md)[name] = value
}

// ReflectData allow to store form values to reflected struct
type ReflectData struct {
	v reflect.Value
}

// NewReflectData create data wrapper on reflected value
func NewReflectData(v reflect.Value) *ReflectData {

	return &ReflectData{v}
}

// Set implements Data interface
func (rd *ReflectData) Set(name string, value interface{}) {
	field := rd.v.FieldByName(name)
	if field.CanSet() {
		field.Set(reflect.ValueOf(value))
	}
}
