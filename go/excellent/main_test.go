package main

import "testing"

func TestEvenOrOdd(t *testing.T) {
	tests := []struct {
		number int
		want   string
	}{
		{number: 2, want: "even"},
		{number: 3, want: "odd"},
		{number: 0, want: "even"},
		{number: -1, want: "odd"},
	}

	for _, tt := range tests {
		got := EvenOrOdd(tt.number)
		if got != tt.want {
			t.Errorf("EvenOrOdd(%d) = %q, want %q", tt.number, got, tt.want)
		}
	}
}
