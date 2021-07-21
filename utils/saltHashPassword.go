package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

func DoPasswordsMatch(password string, otherPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(otherPassword))
	return err == nil
}
