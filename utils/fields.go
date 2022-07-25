package utils

// FieldMapper returns a map with the given values.
func FieldMapper(values map[string]string) (fields map[string]string, err error) {
	if len(values) == 0 {
		return nil, LogError(ErrZeroLength.Error(), LogPrefixInfo, LogLevelInfo)
	}

	for key, value := range values {
		err = Validations(
			ValidateNotEmpty(key),
			ValidateNotEmpty(value),
		)
		if err != nil {
			return nil, err
		}
	}

	return values, nil
}
