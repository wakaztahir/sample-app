package validate

import "regexp"

func Email(email string) ValidationResult {
	result := ValidationResult{
		Field: "email",
		IsValid: false,
		Message: "",
	}

	emailExp := regexp.MustCompile(`[^@]+@[^@]+\.[^@]+`)

	if emailExp.MatchString(email) {
		result.IsValid = true
	} else {
		result.Message = "email is invalid"
	}

	return result
}
