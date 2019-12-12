package main

import (
	"testing"
)

func TestGetAmountOfMatchingPasswords(t *testing.T) {
	type args struct {
		inputs []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Test 1",
			args{[]int{123457}},
			0,
		},
		{"Test 2",
			args{[]int{111111}},
			0,
		},
		{"Test 3",
			args{[]int{123456}},
			0,
		},
		{"Test 4",
			args{[]int{122345}},
			1,
		},
		{"Test 5",
			args{[]int{144452}},
			0,
		},
		{"Test 6",
			args{[]int{222222}},
			0,
		},
		{"Test 7",
			args{[]int{122230}},
			0,
		},
		{"Test 8",
			args{[]int{188000}},
			0,
		},
		{"Test 9",
			args{[]int{112233}},
			1,
		},
		{"Test 10",
			args{[]int{123444}},
			0,
		},
		{"Test 11",
			args{[]int{111122}},
			1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getAmountOfMatchingPasswords(tt.args.inputs); got != tt.want {
				t.Errorf("getAmountOfMatchingPasswords() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hasAdjacentDigits(t *testing.T) {
	type args struct {
		check int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Test 111111", args{111111}, true},
		{"Test 123456", args{123456}, false},
		{"Test 123345", args{123345}, true},
		{"Test 654321", args{654321}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasAdjacentDigits(tt.args.check); got != tt.want {
				t.Errorf("hasAdjacentDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
