package bootstrap

import (
	"html/template"

	"github.com/vporoshok/forms"
)

type BootstrapRenderer struct{}

func New() (*BootstrapRenderer, error) {
	return nil, nil
}

func (renderer *BootstrapRenderer) Render(form *forms.Form) template.HTML {
	return ""
}
