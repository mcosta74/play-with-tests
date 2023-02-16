package utils

import (
	"testing"
	"time"
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReplaceSpaces(tt.args.input); got != tt.want {
				t.Errorf("ReplaceSpaces() = %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCleanup(t *testing.T) {
	t.Cleanup(func() {
		t.Log("First Cleanup")
	})

	t.Cleanup(func() {
		t.Log("Second Cleanup")
	})

	t.Log("setup")

	t.Run("first", func(t *testing.T) {
		t.Log("inside first")
	})

	t.Run("second", func(t *testing.T) {
		t.Log("inside second")
	})

	t.Log("teardown")
}

func TestGroup(t *testing.T) {
	t.Log("main setup")

	t.Run("group", func(t *testing.T) {

		t.Log("group setup")

		tests := []string{"a", "b", "c"}

		for _, tt := range tests {
			tt := tt
			t.Run(tt, func(t *testing.T) {
				t.Parallel()
				t.Logf("start %s", tt)
				defer t.Logf("stop %s", tt)

				time.Sleep(1000 * time.Millisecond)
			})
		}
		t.Log("group teardown")
	})
	t.Log("main teardown")
}
