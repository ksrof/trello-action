/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"errors"
	"log"
)

var (
	ErrZeroLength   error = errors.New("value should not have a length of zero")
	ErrInvalidType  error = errors.New("value type is invalid")
	ErrInvalidMatch error = errors.New("value does not match the regexp pattern")

	LogPrefixInfo string = "[INFO]"

	LogLevelInfo int = 1
)

type Errors func() error

// NewError takes a set of options and returns
// a newly customized error.
func NewError(opts ...Errors) (err error) {
	if len(opts) == 0 {
		return ErrZeroLength
	}

	for _, opt := range opts {
		err = opt()
		return err
	}

	return nil
}

// WithMessage sets the message of the error.
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

// WithLogger sets the message of the error,
// and logs it using the standard logger.
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
		default:
			log.Println(errMessage)
			return err
		}
	}
}
