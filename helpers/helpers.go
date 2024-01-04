package helpers

import "regexp"

func IsValidInput(text string) bool {
	// Use a regular expression to check if the tiinputtle contains only letters and spaces
	match, _ := regexp.MatchString("^[a-zA-Z][a-zA-Z0-9 _-]{2,50}$", text)
	return match
}
