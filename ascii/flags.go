package fs

import "strings"

// SpecialCharacters takes a string input and handles various special characters within it.
// It returns the modified string and a boolean indicating whether a special character was found.
func SpecialCharacters(str string) (string, bool) {
	// Handle newline character
	if str == "\\n" {
		return "", false
	}

	// Handle special characters
	if strings.Contains(str, "\\a") {
		return "Error: Bell Character", true
	}
	if strings.Contains(str, "\\v") {
		return "Error: Vertical tab character", true
	}
	if strings.Contains(str, "\\f") {
		return "Error: Form feed character", true
	}
	if strings.Contains(str, "\\r") {
		return "Error: Carriage return character", true
	}

	// Handle tab characters
	str = strings.ReplaceAll(str, "\\t", "    ")
	// Replace any occurrences of the literal tab character "\t" with 4 spaces

	// Handle backspace tabs
	str = strings.ReplaceAll(str, "\\b", "\b")
	// Replace any occurrences of the backspace character "\b" with the actual backspace character

	// Remove any characters that are preceded by a backspace character
	for {
		index := strings.Index(str, "\b")
		if index == -1 {
			// No more backspace characters found, exit the loop
			break
		}
		if index > 0 {
			// Remove the character before the backspace
			str = str[:index-1] + str[index+1:]
		} else {
			// Remove the backspace character at the beginning of the string
			str = str[index+1:]
		}
	}

	// Return the modified string and a boolean indicating whether a special character was found (false in this case)
	return str, false
}
