package playwithtests

import (
	"runtime"
	"strings"
	"testing"
)

func TestAdd(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 + 0", args{0, 0}, 0},
		{"1 + 0", args{1, 0}, 1},
		{"1 + 1", args{1, 1}, 2},
		{"2 + 1", args{2, 1}, 3},
		{"2 + 2", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"0 * 0", args{0, 0}, 0},
		{"0 * 1", args{0, 1}, 0},
		{"1 * 0", args{1, 0}, 0},
		{"1 * 1", args{1, 1}, 1},
		{"1 * 2", args{1, 2}, 2},
		{"2 * 2", args{2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Mul(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Mul() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFailOn1_20(t *testing.T) {
	version := runtime.Version()
	if strings.HasPrefix(version, "go1.20") {
		t.Fatalf("failing for %s", version)
	}
}
