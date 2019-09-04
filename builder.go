package forms

// Builder helper for quick build forms
type Builder struct {
	Rows []Row
}

// New is a constructor
func New() *Builder {
	return new(Builder)
}

// BuilderOption additional options to fill form row
type BuilderOption interface {
	Apply(*Row)
}

type fnBuilderOption func(*Row)

func (fn fnBuilderOption) Apply(row *Row) {
	fn(row)
}

// Label set label of form row
func (builder *Builder) Label(label string) BuilderOption {
	return fnBuilderOption(func(row *Row) {
		row.Label = label
	})
}

// Tooltip set tooltip of form row
func (builder *Builder) Tooltip(tooltip string) BuilderOption {
	return fnBuilderOption(func(row *Row) {
		row.Tooltip = tooltip
	})
}

// Required set required flag of form row
func (builder *Builder) Required(required bool) BuilderOption {
	return fnBuilderOption(func(row *Row) {
		row.Required = required
	})
}

// Readonly set readonly flag of form row
func (builder *Builder) Readonly(readonly bool) BuilderOption {
	return fnBuilderOption(func(row *Row) {
		row.Readonly = readonly
	})
}

// Disabled set disabled flag of form row
func (builder *Builder) Disabled(disabled bool) BuilderOption {
	return fnBuilderOption(func(row *Row) {
		row.Disabled = disabled
	})
}

// AddControl to form
func (builder *Builder) AddControl(name string, control Control, opts ...BuilderOption) *Builder {
	row := Row{
		Name:    name,
		Label:   name,
		Control: control,
	}
	for _, opt := range opts {
		opt.Apply(&row)
	}
	builder.Rows = append(builder.Rows, row)

	return builder
}

// Build form
func (builder *Builder) Build(data Data) *Form {
	return &Form{
		Rows: builder.Rows,
		Data: data,
	}
}
