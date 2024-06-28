package fs

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

// bannerMap is a map that stores the ASCII art for different banner files
var bannerMap map[string]string

// init initializes the bannerMap and loads the ASCII art from the banner files
func init() {
	// Initialize the bannerMap
	bannerMap = make(map[string]string)

	// Load the ASCII art from the banner files
	loadBanner("standard.txt")
	loadBanner("shadow.txt")
	loadBanner("thinkertoy.txt")
}

// loadBanner reads the contents of a banner file and stores it in the bannerMap
func loadBanner(filename string) {
	// Construct the file path
	filePath := filepath.Join(".", "banner", filename)

	// Open the file
	file, err := os.Open(filePath)
	if err != nil {

		fmt.Printf("Error opening file %s: %v\n", filePath, err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)

	// Initialize an empty slice to store the lines
	var lines []string

	// Read the file line by line
	for scanner.Scan() {
		// Append each line to the lines slice
		lines = append(lines, scanner.Text())
	}

	// Check if the file is empty
	if len(lines) == 0 {
		return
	}

	// Check for any errors while reading the file
	if err := scanner.Err(); err != nil {
		// Handle error reading the file
		fmt.Printf("Error reading file %s: %v\n", filename, err)
		return
	}

	// Join the lines with newline characters and store it in the bannerMap
	bannerMap[filename] = strings.Join(lines, "\n")
}

// GetLetterArray retrieves the ASCII art representation for a given character from the specified banner file
func GetLetterArray(char rune, bannerStyle string) []string {
	// Check if the banner file exists
	banner, ok := bannerMap[bannerStyle]
	if !ok {

		fmt.Println("File doesn't exist")
		os.Exit(1)
	}

	// Split the banner into lines
	alphabet := strings.Split(banner, "\n")

	// Calculate the starting index for the character
	start := (int(char) - 32) * 9

	// Check if the starting index is within the bounds of the alphabet
	if start < 0 || start >= len(alphabet) {
		return []string{}
	}

	// Slice the alphabet to get the ASCII art for the character
	arr := alphabet[start : start+9]

	// Return the ASCII art
	return arr
}

// PrintAscii prints the ASCII art representation of a given string
// PrintAsciiToFile writes the ASCII art representation of a given string to the specified file
func PrintAsciiToFile(str, bannerStyle string, file *os.File) {
	// Split the string into lines
	lines := strings.Split(str, "\n")

	// Initialize an empty slice to store the letters
	letters := [][]string{}

	// Iterate over each line
	for _, line := range lines {
		for _, letter := range line {
			if letter < 32 || letter > 126 {
				fmt.Fprintf(file, "Non-ASCII character '%c' encountered\n", letter)
				return
			}

			// Get the ASCII art for the character
			arr := GetLetterArray(rune(letter), bannerStyle)

			// Append the ASCII art to the letters slice
			letters = append(letters, arr)

			// Check if the character is a newline
			if letter == '\n' {
				// Write a newline to the file
				_, _ = fmt.Fprintln(file)
			}
		}
	}

	// Write the ASCII art to the file vertically
	for i := 1; i < 9; i++ {
		// Iterate over each letter
		for _, letter := range letters {
			// Check if the letter has less than i lines
			if len(letter) < i {
				_, _ = fmt.Fprintln(file, "Error: File content modified")
				return
			}

			// Write the i-th line of the letter to the file
			_, _ = fmt.Fprintf(file, "%s", letter[i-1])
		}

		// Write a newline to the file
		_, _ = fmt.Fprintln(file)
	}
}
