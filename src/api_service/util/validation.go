package util

import (
	"fmt"
	"regexp"
)

// func ValidateURL(url string) error {
// 	// URL pattern regular expression
// 	urlPattern := `^(http(s)?:\/\/.)[-a-zA-Z0-9@:%._\+~#=]{2,256}\.[a-z]{2,6}\b([-a-zA-Z0-9@:%_\+.~#?&//=]*)$`

// 	// Compile the regular expression pattern
// 	r := regexp.MustCompile(urlPattern)

// 	// Check if the URL matches the pattern
// 	if !r.MatchString(url) {
// 		return fmt.Errorf("invalid URL format")
// 	}

// 	return nil
// }

func ValidateEmail(email string) error {
	// Email pattern regular expression
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the regular expression pattern
	r := regexp.MustCompile(emailPattern)

	// Check if the email matches the pattern
	if !r.MatchString(email) {
		return fmt.Errorf("invalid email address")
	}

	return nil
}

func ValidateUsername(username string) error {
	// Username pattern regular expression
	usernamePattern := "^[a-zA-Z0-9_-]+$"

	// Compile the regular expression pattern
	r := regexp.MustCompile(usernamePattern)

	// Check if the username matches the pattern
	if !r.MatchString(username) {
		return fmt.Errorf("invalid username format. Username accepts only alphanumeric characters, dash and underscore")
	}

	return nil
}
