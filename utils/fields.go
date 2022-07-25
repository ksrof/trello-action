/*
Copyright 2022 Kevin Suñer
SPDX-License-Identifier: Apache-2.0
*/

package utils

// FieldMapper returns a map with the given values.
func FieldMapper(values map[string]string) (map[string]string, error) {
	if len(values) == 0 {
		return nil, LogError(ErrZeroLength.Error(), LogPrefixInfo, LogLevelInfo)
	}

	for key, value := range values {
		err := Validations(
			ValidateNotEmpty(key),
			ValidateNotEmpty(value),
		)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
