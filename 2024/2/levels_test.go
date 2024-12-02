package main

import "testing"

func Test_calculateLevels(t *testing.T) {
	type args struct {
		input string
	}
	input := args{`7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9`}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", input, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLevels(tt.args.input); got != tt.want {
				t.Errorf("calculateLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}
