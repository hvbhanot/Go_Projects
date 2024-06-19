package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword takes a password string and generates a bcrypt hash based on it.
// It returns the hashed password as a string and any error that occurred during the generation.
func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	return string(hashedPassword), err
}

// CheckPasswordHash checks if a plain text password matches a hashed password.
// It uses the bcrypt.CompareHashAndPassword function to compare the hashed password
// with the plain text password. If the comparison is successful, it returns true.
// Otherwise, it returns false.
func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
