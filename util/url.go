package util

import (
	"fmt"
	"net/url"
)

func IsValidURL(input string) (bool, error) {
	u, err := url.Parse(input)
	if err != nil {
		return false, fmt.Errorf("invalid URL: %v", err)
	}
	if u.Scheme == "" {
		return false, fmt.Errorf("invalid URL: URL scheme is missing, it should be http or https")
	}
	if u.Scheme != "http" && u.Scheme != "https" {
		return false, fmt.Errorf("invalid URL: URL scheme should be http or https")
	}
	if u.Host == "" {
		return false, fmt.Errorf("invalid URL: '%v' is not a valid web URL host", u.Host)
	}
	return true, nil
}
