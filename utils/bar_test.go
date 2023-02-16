package utils

import "testing"

func TestReplaceSpaces(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"empty string", args{""}, ""},
		{"string without spaces", args{"Something"}, "Something"},
		{"string with spaces", args{"this is very interesting"}, "this_is_very_interesting"},
		{"skip this", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "skip" {
				t.Skip("skipped test")
			}

			if got := ReplaceSpaces(tt.args.input); got != tt.want {
				t.Errorf("ReplaceSpaces() = %q, want %q", got, tt.want)
			}
		})
	}
}
