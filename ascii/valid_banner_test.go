package fs

import "testing"

func TestIsValidBanner(t *testing.T) {
	type args struct {
		bannerStyle string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Valid banner style",
			args: args{"standard"},
			want: true,
		},
		{
			name: "Invalid banner style",
			args: args{"invalid"},
			want: false,
		},
		{
			name: "Banner style with .txt",
			args: args{"standard.txt"},
			want: false,
		},
		{
			name: "Banner style without .txt",
			args: args{"standard"},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValidBanner(tt.args.bannerStyle); got != tt.want {
				t.Errorf("IsValidBanner() = %v, want %v", got, tt.want)
			}
		})
	}
}
