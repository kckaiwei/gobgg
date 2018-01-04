package gobgg

import "fmt"

// StatusError is an error returned when there's HTTP Response error
type StatusError struct {
	message 	string
	code		int
}

func (se *StatusError) Error() string {
	return fmt.Sprintf("StatusError: %v", se.code)
}