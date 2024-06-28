package fs

import (
	"testing"
)

func TestSpecialCharacters(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantErr bool
	}{
		{"Empty string", "", "", false},
		{"Newline character", "\\n", "", false},
		{"Bell character", "Hello\\aWorld", "Error: Bell Character", true},
		{"Vertical tab character", "Hello\\vWorld", "Error: Vertical tab character", true},
		{"Form feed character", "Hello\\fWorld", "Error: Form feed character", true},
		{"Carriage return character", "Hello\\rWorld", "Error: Carriage return character", true},
		{"Tab character", "Hello\\tWorld", "Hello    World", false},
		{"Backspace tab", "Hello\\bWorld", "HellWorld", false},
		{"Multiple special characters", "Hello\\a\\v\\f\\rWorld", "Error: Bell Character", true},
		{"Mixed characters", "Hello\\tWorld\\n\\b", "Hello    World\\", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := SpecialCharacters(tt.input)
			if gotErr != tt.wantErr {
				t.Errorf("SpecialCharacters() error = %v, wantErr %v", gotErr, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SpecialCharacters() = %v, want %v", got, tt.want)
			}
		})
	}
}
