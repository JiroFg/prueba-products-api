package exceptions

import "fmt"

type NotFoundError struct {
	ElementName string
	ElementId   string
}

func (e *NotFoundError) Error() string {
	return fmt.Sprintf("%s with ID %s not found", e.ElementName, e.ElementId)
}
