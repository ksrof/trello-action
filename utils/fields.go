/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils

type Field func() (interface{}, error)

// NewFields takes a set of field options and returns
// the fields or an error in case of failure.
func NewFields(opts ...Field) (fields interface{}, err error) {
	if len(opts) == 0 {
		return nil, NewError(
			WithLogger(
				ErrZeroLength.Error(),
				LogPrefixInfo,
				LogLevelInfo,
			),
		)
	}

	for _, opt := range opts {
		fields, err = opt()
		if err != nil {
			return nil, NewError(
				WithLogger(
					err.Error(),
					LogPrefixInfo,
					LogLevelInfo,
				),
			)
		}
	}

	return fields, nil
}

// WithMap validates the keys and the values of a given map
// and returns it if there is no error.
func WithMap(fields map[string]string) Field {
	return func() (interface{}, error) {
		if len(fields) == 0 {
			return nil, ErrZeroLength
		}

		for key, value := range fields {
			err := Validations(
				ValidateNotEmpty(key),
				ValidateNotEmpty(value),
			)
			if err != nil {
				return nil, err
			}
		}

		return fields, nil
	}
}
