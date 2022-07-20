/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"errors"
	"log"
	"regexp"
	"strconv"
)

var (
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

// ValidateNotEmpty checks that the length of the
// given value isn't zero.
func ValidateNotEmpty[T any](value T) Validation {
	return func() error {
		switch any(value).(type) {
		case string:
			if len(any(value).(string)) == 0 {
				log.Printf(
					// TODO: Print the value.
					"failed to validate any(value).(string), error: %s\n",
					ErrEmptyValue.Error(),
				)
				return ErrEmptyValue
			}

			return nil
		default:
			return ErrInvalidType
		}
	}
}

// ValidateRegexp checks that the given pattern
// matches the given value.
func ValidateRegexp[T any](pattern regexp.Regexp, value T) Validation {
	return func() error {
		switch any(value).(type) {
		case string:
			ok := pattern.MatchString(any(value).(string))
			if !ok {
				log.Printf(
					"failed to match any(value).(string), error: %s\n",
					ErrInvalidMatch.Error(),
				)
				return ErrInvalidMatch
			}

			return nil
		case int:
			strVal := strconv.Itoa(any(value).(int))
			ok := pattern.Match([]byte(strVal))
			if !ok {
				log.Printf(
					"failed to match any(value).(int), error: %s\n",
					ErrInvalidMatch.Error(),
				)

				return ErrInvalidMatch
			}

			return nil
		default:
			return ErrInvalidType
		}
	}
}
