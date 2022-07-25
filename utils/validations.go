/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils

import (
	"regexp"
	"strconv"
)

type Validation func() error

// Validations takes a set of validations and returns
// an error in case of failure.
func Validations(opts ...Validation) error {
	if len(opts) == 0 {
		return LogError(ErrZeroLength.Error(), LogPrefixInfo, LogLevelInfo)
	}

	for _, opt := range opts {
		err := opt()
		if err != nil {
			return LogError(err.Error(), LogPrefixInfo, LogLevelInfo)
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
				return ErrZeroLength
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
				return ErrInvalidMatch
			}

			return nil
		case int:
			strVal := strconv.Itoa(any(value).(int))
			ok := pattern.Match([]byte(strVal))
			if !ok {
				return ErrInvalidMatch
			}

			return nil
		default:
			return ErrInvalidType
		}
	}
}
