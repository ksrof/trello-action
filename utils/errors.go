package utils

import (
	"errors"
	"log"
)

var (
	// Errors.
	ErrZeroLength   error = errors.New("value length can not be zero")
	ErrInvalidMatch error = errors.New("value does not match the regexp pattern")
	ErrInvalidType  error = errors.New("value type is not valid")

	// Prefixes.
	LogPrefixInfo string = "[INFO]"

	// Levels.
	LogLevelInfo int = 1
)

// LogError returns a newly created error with the given message
// and logs it using the standard logger.
func LogError(errStr, prefix string, level int) (err error) {
	err = Validations(
		ValidateNotEmpty(errStr),
	)
	if err != nil {
		return err
	}

	log.SetPrefix(prefix)
	err = errors.New(errStr)

	switch level {
	case 1:
		log.Println(err)
		return err
	default:
		log.Println(err)
		return err
	}
}
