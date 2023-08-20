package util

import "golang.org/x/crypto/bcrypt"

func GeneratePassword() (string, string) {
	password := GenerateRandomString(20)
	hashedPassword := HashPassword(password)
	return password, hashedPassword
}

func HashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword)
}

func VerifyPassword(hashedPassword string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}
