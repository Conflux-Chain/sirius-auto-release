package utils

import (
	"fmt"
	"strconv"
)

func ValidateNumber(fieldName string) func(string) error {
	return func(s string) error {
		if s == "" {
			return fmt.Errorf("%s cannot be empty", fieldName)
		}
		if _, err := strconv.Atoi(s); err != nil {
			return fmt.Errorf("%s must be a number", fieldName)
		}
		return nil
	}
}
