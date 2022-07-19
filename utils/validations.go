package utils

import (
	"errors"
	"log"
)

var (
	ErrNilValue    error = errors.New("value is nil")
	ErrInvalidType error = errors.New("invalid type provided")
	ErrUnknown     error = errors.New("unknown error")
)

func ValidType[T any](value T) error {
	if any(value) == nil {
		log.Printf(
			"failed type assertion any(value), error: %s",
			ErrNilValue.Error(),
		)
		return ErrNilValue
	}

	switch any(value) {
	case any(value).(string):
		log.Println("value is string type")
		return nil
	case any(value).(int):
		log.Println("value is integer type")
		return nil
	}

	return ErrUnknown
}
