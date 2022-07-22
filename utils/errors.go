package utils

import (
	"errors"
	"log"
)

// TODO: Test that every fn works as expected.
// TODO: Refactor and implement this error handler.

type Errors func() error

func NewError(opts ...Errors) (err error) {
	if len(opts) == 0 {
		return ErrEmptyOptions
	}

	for _, opt := range opts {
		err = opt()
		return err
	}

	return nil
}

func WithMessage(errMessage string) Errors {
	return func() error {
		err := Validations(
			ValidateNotEmpty(errMessage),
		)
		if err != nil {
			return err
		}

		err = errors.New(errMessage)
		return err
	}
}

func WithLogger(errMessage, prefix string, level int) Errors {
	return func() error {
		err := Validations(
			ValidateNotEmpty(errMessage),
		)
		if err != nil {
			return err
		}

		log.SetPrefix(prefix)
		err = errors.New(errMessage)

		switch level {
		case 1:
			log.Println(errMessage)
			return err
		case 2:
			log.Fatalln(errMessage)
			return err
		case 3:
			log.Panicln(errMessage)
			return err
		default:
			log.Println(errMessage)
			return err
		}
	}
}
