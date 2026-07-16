package exceptions

import "fmt"

type FieldError struct {
	FieldName string
	Message   string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("Campo: \"%s\" con error: \"%s\"", e.FieldName, e.Message)
}
