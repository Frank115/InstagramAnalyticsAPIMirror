package domain

import (
	"fmt"
)

type ApiError struct {
	Status  int
	Err     error
	Message string
}

func (a ApiError) Error() string {
	return fmt.Sprintf("El error es: %v", a.Err.Error())
}
