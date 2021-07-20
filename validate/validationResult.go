package validate

type ValidationResult struct {
	Field   string `json:"field"`
	IsValid bool   `json:"is_valid"`
	Message string `json:"message"`
}
