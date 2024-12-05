package main

import (
	"testing"
)

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
	input2 := args{`36 34 31 30 29 28 27
53 51 49 46 43
95 92 90 88 85 83
74 73 72 70 67 64
23 20 17 16 14 13
95 92 89 87 85 82 81 79
36 35 34 32 31 30 29
50 52 54 57 60 63 64
58 61 62 63 64 65 68 70
39 40 42 43 45 48 51
85 84 81 79 77 74
45 46 47 49 52
15 16 17 20 21 22
86 83 80 78 76 73 70 68
51 54 57 59 62
71 73 76 77 79 80
68 69 72 74 77 78 79 81
13 11 8 6 5 3
50 52 53 56 57 59 60
60 62 64 66 69
80 83 84 86 88`}
	input3 := args{`40 42 38 36 33 32 29 29
29 31 29 28 25 24 20 16
80 82 79 76 75 71 69 63
49 52 50 47 46 41 39
78 79 78 77 72 69 66 67
12 13 11 5 5
12 15 14 9 8 7 6 2
44 46 45 43 38 36 30
61 61 58 56 53 52 50 47
35 35 32 31 28 30
46 46 44 42 42
80 80 77 76 72
63 63 60 59 56 55 49`}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"sample", input, 2},
		{"success", input2, 21},
		{"fail", input3, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateLevels(tt.args.input); got != tt.want {
				t.Errorf("calculateLevels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_safeAdjacent(t *testing.T) {
	type args struct {
		l string
	}

	isSafeInput1 := "1 3 6 7 9"
	isSafeInput2 := "30 33 34 35 38"
	isSafeInput3 := "49 51 52 54 56"
	isSafeInput4 := "22 24 25 28 30"
	isSafeInput5 := "18 19 22 24 27 30 33 35"
	isSafeInput6 := "83 84 87 89 92"
	isSafeInput7 := "32 35 36 37 40 43"
	isSafeInput8 := "18 17 16 14 13 10 9"
	isSafeInput9 := "53 51 49 46 45 42"
	isSafeInput10 := "64 65 68 70 72"
	isSafeInput11 := "72 69 68 65 64"
	isSafeInput12 := "55 53 50 49 46 43 42 39"
	isSafeInput13 := "53 52 50 48 45 42 39"
	isSafeInput14 := "71 73 75 78 81"
	isSafeInput15 := "83 85 86 89 91 94"
	isSafeInput16 := "84 85 87 89 92 93"
	isSafeInput17 := "27 24 23 22 19 18"
	isSafeInput18 := "19 22 24 26 27"
	isSafeInput19 := "58 57 56 55 52 49 47"
	isSafeInput20 := "80 79 76 75 74"
	isSafeInput21 := "89 86 83 80 77 74"

	isUnsafeInput1 := "8 6 4 4 1"
	isUnsafeInput2 := "1 2 7 8 9"
	isUnsafeInput3 := "9 7 6 2 1"
	isUnsafeInput4 := "43 46 47 51 52 53 56 60"
	isUnsafeInput5 := "13 14 17 20 21 26 26"
	isUnsafeInput6 := "96 97 96 93 89 88 86"
	isUnsafeInput7 := "39 45 46 43 44 45"
	isUnsafeInput8 := "87 80 76 75 72 72"
	isUnsafeInput10 := "67 64 61 58 57 53"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Safe test1", args{isSafeInput1}, true},
		{"Safe test2", args{isSafeInput2}, true},
		{"Safe test3", args{isSafeInput3}, true},
		{"Safe test4", args{isSafeInput4}, true},
		{"Safe test5", args{isSafeInput5}, true},
		{"Safe test6", args{isSafeInput6}, true},
		{"Safe test7", args{isSafeInput7}, true},
		{"Safe test8", args{isSafeInput8}, true},
		{"Safe test9", args{isSafeInput9}, true},
		{"Safe test10", args{isSafeInput10}, true},
		{"Safe test11", args{isSafeInput11}, true},
		{"Safe test12", args{isSafeInput12}, true},
		{"Safe test13", args{isSafeInput13}, true},
		{"Safe test14", args{isSafeInput14}, true},
		{"Safe test15", args{isSafeInput15}, true},
		{"Safe test16", args{isSafeInput16}, true},
		{"Safe test17", args{isSafeInput17}, true},
		{"Safe test18", args{isSafeInput18}, true},
		{"Safe test19", args{isSafeInput19}, true},
		{"Safe test20", args{isSafeInput20}, true},
		{"Safe test21", args{isSafeInput21}, true},

		{"Unsafe test1", args{isUnsafeInput1}, false},
		{"Unsafe test2", args{isUnsafeInput2}, false},
		{"Unsafe test3", args{isUnsafeInput3}, false},
		{"Unsafe test4", args{isUnsafeInput4}, false},
		{"Unsafe test5", args{isUnsafeInput5}, false},
		{"Unsafe test6", args{isUnsafeInput6}, false},
		{"Unsafe test7", args{isUnsafeInput7}, false},
		{"Unsafe test8", args{isUnsafeInput8}, false},
		{"Unsafe test10", args{isUnsafeInput10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := safeAdjacent(tt.args.l); got != tt.want {
				t.Errorf("safeAdjacent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_uniDirectional(t *testing.T) {
	type args struct {
		l string
	}
	isUnidirectionalAsc := args{"2 3 5 6 8"}
	isUnidirectionalDesc := args{"7 6 4 2 1"}
	isNotUnidirectional1 := args{"1 3 2 4 5"}
	isNotUnidirectional2 := args{"8 6 4 4 1"}
	isNotUnidirectional3 := args{"8 5 3 2 3"}

	isUnsafeInput5 := "13 14 17 20 21 26 26"
	isUnsafeInput6 := "96 97 96 93 89 88 86"
	isUnsafeInput7 := "39 45 46 43 44 45"
	isUnsafeInput8 := "87 80 76 75 72 72"
	isUnsafeInput9 := "86 88 89 90 92 93 92 91"
	isSafeInput1 := "1 3 6 7 9"
	isSafeInput2 := "30 33 34 35 38"
	isSafeInput3 := "49 51 52 54 56"
	isSafeInput4 := "22 24 25 28 30"
	isSafeInput5 := "18 19 22 24 27 30 33 35"
	isSafeInput6 := "83 84 87 89 92"
	isSafeInput7 := "32 35 36 37 40 43"
	isSafeInput8 := "18 17 16 14 13 10 9"
	isSafeInput9 := "53 51 49 46 45 42"
	isSafeInput10 := "64 65 68 70 72"
	isSafeInput11 := "72 69 68 65 64"
	isSafeInput12 := "55 53 50 49 46 43 42 39"
	isSafeInput13 := "53 52 50 48 45 42 39"
	isSafeInput14 := "71 73 75 78 81"
	isSafeInput15 := "83 85 86 89 91 94"
	isSafeInput16 := "84 85 87 89 92 93"
	isSafeInput17 := "27 24 23 22 19 18"
	isSafeInput18 := "19 22 24 26 27"
	isSafeInput19 := "58 57 56 55 52 49 47"
	isSafeInput20 := "80 79 76 75 74"
	isSafeInput21 := "89 86 83 80 77 74"

	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Safe Ascending", isUnidirectionalAsc, true},
		{"Safe Descending", isUnidirectionalDesc, true},
		{"Unsafe", isNotUnidirectional1, false},
		{"Unsafe 2", isNotUnidirectional2, false},
		{"Unsafe test10", isNotUnidirectional3, false},

		{"Unsafe test5", args{isUnsafeInput5}, false},
		{"Unsafe test6", args{isUnsafeInput6}, false},
		{"Unsafe test7", args{isUnsafeInput7}, false},
		{"Unsafe test8", args{isUnsafeInput8}, false},
		{"Unsafe test9", args{isUnsafeInput9}, false},

		{"Safe test1", args{isSafeInput1}, true},
		{"Safe test2", args{isSafeInput2}, true},
		{"Safe test3", args{isSafeInput3}, true},
		{"Safe test4", args{isSafeInput4}, true},
		{"Safe test5", args{isSafeInput5}, true},
		{"Safe test6", args{isSafeInput6}, true},
		{"Safe test7", args{isSafeInput7}, true},
		{"Safe test8", args{isSafeInput8}, true},
		{"Safe test9", args{isSafeInput9}, true},
		{"Safe test10", args{isSafeInput10}, true},
		{"Safe test11", args{isSafeInput11}, true},
		{"Safe test12", args{isSafeInput12}, true},
		{"Safe test13", args{isSafeInput13}, true},
		{"Safe test14", args{isSafeInput14}, true},
		{"Safe test15", args{isSafeInput15}, true},
		{"Safe test16", args{isSafeInput16}, true},
		{"Safe test17", args{isSafeInput17}, true},
		{"Safe test18", args{isSafeInput18}, true},
		{"Safe test19", args{isSafeInput19}, true},
		{"Safe test20", args{isSafeInput20}, true},
		{"Safe test21", args{isSafeInput21}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniDirectional(tt.args.l); got != tt.want {
				t.Errorf("uniDirectional() = %v, want %v", got, tt.want)
			}
		})
	}
}
