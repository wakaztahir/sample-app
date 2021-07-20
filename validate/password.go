package validate

func Password(password string) ValidationResult {
	result := ValidationResult{
		Field:   "password",
		IsValid: false,
		Message: "",
	}

	if len(password) > 4 && len(password) <= 70 {
		result.IsValid = true
	} else {
		result.Message = "password must be 5-70 characters long"
	}

	return result
}
