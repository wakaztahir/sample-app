package validate

import "regexp"

func Name(name string) ValidationResult {
	//Checking Length Of Name
	result := ValidationResult{
		Field: "name",
		IsValid: false,
		Message: "",
	}

	if len(name) > 2 && len(name) <= 70 {
		//Succeeded Length Check
		nameExp := regexp.MustCompile(`[^a-zA-Z. ]`)

		if !nameExp.MatchString(name) {
			result.IsValid = true
		} else {
			result.Message = "name must only contain alphabetical characters and spaces"
		}

	} else {
		result.Message = "name length must be 2-70 characters long"
	}

	return result
}
