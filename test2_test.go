package main

import "testing"

func TestDecode(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "LLRR=",
			want:  "210122",
		},
		{
			input: "==RLL",
			want:  "000210",
		},
		{
			input: "=LLRR",
			want:  "221012",
		},
		{
			input: "RRL=R",
			want:  "012001",
		},
	}
	for _, tt := range tests {
		t.Run("Decode", func(t *testing.T) {
			if got := Decode(tt.input); got != tt.want {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}
