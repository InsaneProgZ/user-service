package model

import "fmt"

type UserError struct {
	Message string
}

func (e *UserError) Error() string {
	return fmt.Sprint(e.Message)
}
