package utils

import (
	"runtime"
	"strings"
	"testing"
)

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
		{"skip", args{""}, ""},
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

func TestFoo(t *testing.T) {
	t.Skip("something to ignore")
}

func TestBar(t *testing.T) {
	t.Run("hello", func(t *testing.T) {
		t.Skip("something to ignore")
	})

	t.Run("foo", func(t *testing.T) {
		// t.Skip("something to ignore")
	})
}

func TestBaz(t *testing.T) {

	tests := []struct {
		name string
		in   int
		out  int
	}{
		{"a", 1, 2},
		{"b", 2, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Skipf("skipping %d", tt.in)

		})
	}
}

func TestFailOn1_19(t *testing.T) {
	version := runtime.Version()
	if strings.HasPrefix(version, "go1.19") {
		t.Fatalf("failing for %s", version)
	}
}
