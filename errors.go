package forms

type constantError string

func (ce constantError) Error() string {
	return string(ce)
}

const (
	ErrInvalidDataType constantError = "data should be a pointer to struct"
)
