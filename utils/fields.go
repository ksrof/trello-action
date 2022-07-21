/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils

import "log"

type Field func() (interface{}, error)

// NewFieldsMapper takes a set of field options and returns
// the fields or an error in case of failure.
func NewFieldsMapper(opts ...Field) (fields interface{}, err error) {
	if len(opts) == 0 {
		return nil, ErrEmptyOptions
	}

	for _, opt := range opts {
		fields, err = opt()
		if err != nil {
			log.Printf(
				"failed to add field options, error: %s\n",
				err.Error(),
			)
			return nil, err
		}
	}

	return fields, nil
}

// WithMap validates the keys and the values of a given map
// and returns it if there is no error.
func WithMap(fields map[string]string) Field {
	return func() (interface{}, error) {
		for key, value := range fields {
			err := Validations(
				ValidateNotEmpty(key),
				ValidateNotEmpty(value),
			)
			if err != nil {
				log.Printf(
					"failed to validate map keys or values, error: %s\n",
					err.Error(),
				)
				return nil, err
			}
		}

		return fields, nil
	}
}
