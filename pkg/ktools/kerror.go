package ktools

import (
	"errors"
)

// MergeError merges multiple errors into one
func MergeError(errs ...error) error {
	message := ""
	for _, err := range errs {
		if err != nil {
			message += err.Error() + " | "
		}
	}
	if message != "" {
		message = message[:len(message)-3]
		return errors.New(message)
	}
	return nil
}