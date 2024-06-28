package main

import (
	"fmt"
	"os"
	"strings"

	fs "fs/ascii" // Import the ascii package from the local directory
)

func main() {
	// Check if the command-line arguments are valid
	if len(os.Args) < 2 || len(os.Args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\nEX: go run . --output=output.txt something standard")
		return
	}

	// Get the output file, input string, and banner style from the command-line arguments
	var outputFile string
	str := os.Args[1]
	bannerStyle := "standard.txt" // Default banner style

	if len(os.Args) >= 3 {
		if strings.HasPrefix(os.Args[1], "--output=") {
			outputFile = os.Args[1][9:]
			str = os.Args[2]
			if len(os.Args) == 4 {
				bannerStyle = os.Args[3]
			}
		} else {
			str = os.Args[1]
			if len(os.Args) == 3 {
				bannerStyle = os.Args[2]
			}
		}
	}

	// Check if the banner style is a valid file
	if !fs.IsValidBanner(bannerStyle) {
		fmt.Println("Error: Banner style invalid")
		return
	}

	// Check if the input string is empty
	if str == "" {
		return
	}

	// Handle special characters in the input string
	modifiedStr, hasError := fs.SpecialCharacters(str)
	if hasError {
		fmt.Println(modifiedStr)
		return
	}

	// Replace newline characters with "\\n" and split the string into lines
	str = modifiedStr
	str = strings.ReplaceAll(str, "\n", "\\n")
	lines := strings.Split(str, "\\n")

	// Print the ASCII art for each line
	newlineCount := 0
	if outputFile != "" {
		// Open the output file
		file, err := os.Create(outputFile)
		if err != nil {
			fmt.Printf("Error creating file %s: %v\n", outputFile, err)
			return
		}
		defer file.Close()

		// Write the ASCII art to the file
		for _, line := range lines {
			if line == "" {
				newlineCount++
				if newlineCount > 0 {
					_, _ = fmt.Fprintln(file)
				}
			} else {
				fs.PrintAsciiToFile(line, bannerStyle, file) // Call the PrintAsciiToFile function to write to the file
			}
		}
	}
}
