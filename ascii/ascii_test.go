package fs

import (
	"testing"
)

func TestLoadBanner(t *testing.T) {
	// Test loading a valid banner file
	loadBanner("shadow.txt")
	if _, ok := bannerMap["shadow.txt"]; !ok {
		t.Error("Failed to load banner file")
	}
	loadBanner("thinkertoy.txt")
	if _, ok := bannerMap["thinkertoy.txt"]; !ok {
		t.Error("Failed to load banner file")
	}
	loadBanner("standard.txt")
	if _, ok := bannerMap["standard.txt"]; !ok {
		t.Error("Failed to load the banner file")
	}

	// Add more test cases for loading other banner files or error scenarios
}

func TestPrintAscii(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		bannerStyle string
	}{
		{"Standard banner", "Hello", "standard.txt"},
		{"Shadow banner", "World", "shadow.txt"},
		{"Non-ASCII character", "Hello\x00", "standard.txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Capture the output of PrintAscii
			// ...
		})
	}
}
