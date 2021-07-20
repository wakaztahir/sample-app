package validate

import "regexp"

func Username(username string) ValidationResult {
	//Checking Length Of Username
	result := ValidationResult{
		IsValid: false,
		Message: "",
	}

	if len(username) > 4 && len(username) <= 60 {
		//Succeeded Length Check
		nameExp := regexp.MustCompile(`[^a-zA-Z0-9.$]`)

		if !nameExp.MatchString(username) {
			result.IsValid = true
		} else {
			result.Message = "username must only contain alphanumeric characters and no spaces"
		}

	} else {
		result.Message = "username length must be 4-60 characters long"
	}

	return result
}
