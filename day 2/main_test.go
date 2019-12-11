package main

import (
	"reflect"
	"strings"
	"testing"
)

func TestParseIntcode(t *testing.T) {
	type args struct {
		intCode string
		noun    string
		verb    string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"Example case 1", args{"1,0,0,0,99", "0", "0"}, strings.Split("2,0,0,0,99", ",")},
		{"Example case 2", args{"2,3,0,3,99", "3", "0"}, strings.Split("2,3,0,6,99", ",")},
		{"Example case 3", args{"2,4,4,5,99,0", "4", "4"}, strings.Split("2,4,4,5,99,9801", ",")},
		{"Example case 4", args{"1,1,1,4,99,5,6,0,99", "1", "1"}, strings.Split("30,1,1,4,2,5,6,0,99", ",")},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseIntcode(tt.args.intCode, tt.args.noun, tt.args.verb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseIntcode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
