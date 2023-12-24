package functions

import "regexp"

func IsInteger(value string) bool {
	validInteger := regexp.MustCompile(`[0-9]`)
	return validInteger.MatchString(value)
}

func IsOperator(value string) bool {
	validOperator := regexp.MustCompile(`[\+\-\*\/]`)
	return validOperator.MatchString(value)
}
