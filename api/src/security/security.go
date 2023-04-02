package security

import "golang.org/x/crypto/bcrypt"

// Hash receive a string and put a hash
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

// CheckPassword is a method to validate a password
func CheckPassword(stringPassword, hashPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(stringPassword))
}
