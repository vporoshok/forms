package forms_test

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/vporoshok/forms"
)

func TestMapData(t *testing.T) {
	md := make(forms.MapData)
	md.Set("Foo", "Bar")
	assert.Equal(t, "Bar", md["Foo"])
}

func TestReflectData(t *testing.T) {
	data := struct {
		Foo string
	}{}
	rd := forms.NewReflectData(reflect.ValueOf(&data).Elem())
	rd.Set("Foo", "Bar")
	assert.Equal(t, "Bar", data.Foo)
}
