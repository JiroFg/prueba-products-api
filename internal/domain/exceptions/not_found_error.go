package exceptions

import "fmt"

type NotFoundError struct {
	ElementName string
	ElementId   string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s con ID %s no encontrado", e.ElementName, e.ElementId)
}
