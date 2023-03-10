package utils

import "testing"

func TestHello(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Massimo", args{"Massimo"}, "Hello Massimo"},
		{"Marianna", args{"Marianna"}, "Hello Marianna"},
		{"no name", args{""}, "Who are you?"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Hello(tt.args.name); got != tt.want {
				t.Errorf("Hello() = %q, want %q", got, tt.want)
			}
		})
	}
}
