package main

import "testing"

func Test_compareLists(t *testing.T) {
	type args struct {
		a []int
		b []int
	}

	sample_input := args{[]int{3, 4, 2, 1, 3, 3}, []int{4, 3, 5, 3, 9, 3}}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", sample_input, 11}, // TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := compareLists(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("compareLists() = %v, want %v", got, tt.want)
			}
		})
	}
}
