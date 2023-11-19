package dto

import (
	"errors"
	"strings"
)

// ValidateLoop is a function that pass a slice of validation functions and execute them in order
func ValidateLoop(orderExec []func() error) error {
	errMsg := ""
	for _, val := range orderExec {
		if err := val(); err != nil {
			errMsg += err.Error() + " | "
		}
	}
	if errMsg != "" {
		errMsg = strings.TrimSuffix(errMsg, " | ")
		return errors.New(errMsg)
	}
	return nil
}