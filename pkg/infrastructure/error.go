package infrastructure

import "fmt"

type DbError struct {
	message string
	originalError error
}

func (e *DbError) Error() string {
	return fmt.Sprintf("%s:%s", e.message, e.originalError.Error())
}
