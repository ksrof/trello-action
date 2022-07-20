/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

// TODO: Add tests to check that every fn and method work.

package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

var (
	ErrNilValue     error = errors.New("value is nil")
	ErrInvalidType  error = errors.New("invalid type provided")
	ErrEmptyValue   error = errors.New("value is empty")
	ErrInvalidMatch error = errors.New("value doesn't match regexp")
	ErrUnknown      error = errors.New("unknown error")
)

type Validation func() error

// Validations takes a set of validations and returns
// an error in case of failure.
func Validations(validations ...Validation) error {
	for _, validation := range validations {
		err := validation()
		if err != nil {
			log.Printf(
				"failed to add validations, error: %s",
				err.Error(),
			)
			return err
		}
	}

	return nil
}

// ValidateType asserts the type of the given value.
func ValidateType[T any](value T) Validation {
	return func() error {
		if any(value) == nil {
			log.Printf(
				"failed type assertion any(value), error: %s\n",
				ErrNilValue.Error(),
			)
			return ErrNilValue
		}

		switch any(value) {
		case any(value).(string):
			return nil
		case any(value).(int):
			return nil
		}

		return ErrUnknown
	}
}

// ValidateNotZero checks the length of the given value isn't zero.
func ValidateNotZero[T any](value T) Validation {
	return func() error {
		if any(value) == nil {
			log.Printf(
				"failed type assertion any(value), error: %s\n",
				ErrNilValue.Error(),
			)
			return ErrNilValue
		}

		switch any(value) {
		case any(value).(string):
			if len(any(value).(string)) == 0 {
				log.Printf(
					"failed to validate any(value).(string), error: %s\n",
					ErrEmptyValue.Error(),
				)
				return ErrEmptyValue
			}

			return nil
		case any(value).(int):
			if any(value).(int) == 0 {
				log.Printf(
					"failed to validate any(value).(int), error: %s\n",
					ErrEmptyValue.Error(),
				)
				return ErrEmptyValue
			}

			return nil
		}

		return ErrUnknown
	}
}

// ValidateRegexp checks that the given pattern
// matches the given value.
func ValidateRegexp[T any](pattern regexp.Regexp, value T) Validation {
	return func() error {
		if any(value) == nil {
			log.Printf(
				"failed type assertion any(value), error: %s\n",
				ErrNilValue.Error(),
			)
			return ErrNilValue
		}

		switch any(value) {
		case any(value).(string):
			ok := pattern.MatchString(any(value).(string))
			if !ok {
				log.Printf(
					"failed to match any(value).(string), error: %s\n",
					ErrInvalidMatch.Error(),
				)
				return ErrInvalidMatch
			}

			return nil
		case any(value).(int):
			strVal := strconv.Itoa(any(value).(int))
			ok := pattern.Match([]byte(strVal))
			if !ok {
				log.Printf(
					"failed to match any(value).(string), error: %s\n",
					ErrInvalidMatch.Error(),
				)

				return ErrInvalidMatch
			}

			return nil
		}

		return ErrUnknown
	}
}
