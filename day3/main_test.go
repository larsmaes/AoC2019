package main

import (
	"strings"
	"testing"
)

func TestFindClosestIntersectionDistance(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name         string
		args         args
		wantDistance int
	}{
		{"Test 1", args{
			[][]string{
				strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ","),
				strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ","),
			},
		},
			159,
		},
		{"Test 2", args{
			[][]string{
				strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ","),
				strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ","),
			},
		},
			135,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotDistance := findClosestIntersectionDistance(tt.args.paths); gotDistance != tt.wantDistance {
				t.Errorf("findClosestIntersectionDistance() = %v, want %v", gotDistance, tt.wantDistance)
			}
		})
	}
}

func TestFindLowestIntersectionStepCount(t *testing.T) {
	type args struct {
		paths [][]string
	}
	tests := []struct {
		name                string
		args                args
		wantLowestStepcount int
	}{
		{"Test 1", args{
			[][]string{
				strings.Split("R75,D30,R83,U83,L12,D49,R71,U7,L72", ","),
				strings.Split("U62,R66,U55,R34,D71,R55,D58,R83", ",")},
		},
			610},
		{"Test 2", args{
			[][]string{
				strings.Split("R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", ","),
				strings.Split("U98,R91,D20,R16,D67,R40,U7,R15,U6,R7", ",")},
		},
			410},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotLowestStepcount := findLowestIntersectionStepCount(tt.args.paths); gotLowestStepcount != tt.wantLowestStepcount {
				t.Errorf("findLowestIntersectionStepCount() = %v, want %v", gotLowestStepcount, tt.wantLowestStepcount)
			}
		})
	}
}
