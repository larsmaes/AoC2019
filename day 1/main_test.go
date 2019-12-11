package main

import (
	"testing"
)

func Test_calcNeededFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Mass of 12 should yield 2", args{12}, 2},
		{"Mass of 14 should yield 2", args{14}, 2},
		{"Mass of 1969 should yield 654", args{1969}, 654},
		{"Mass of 100756 should yield 33583", args{100756}, 33583},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcNeededFuel(tt.args.mass); got != tt.want {
				t.Errorf("calcNeededFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcTotalModuleFuel(t *testing.T) {
	type args struct {
		mass int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"Mass of 12 should yield 2", args{12}, 2},
		{"Mass of 14 should yield 2", args{14}, 2},
		{"Mass of 1969 should yield 966", args{1969}, 966},
		{"Mass of 100756 should yield 50346", args{100756}, 50346},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcTotalModuleFuel(tt.args.mass); got != tt.want {
				t.Errorf("calcTotalModuleFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Main test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
