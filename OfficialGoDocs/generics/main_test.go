package main

import (
	"reflect"
	"testing"
)

func TestSumInts(t *testing.T) {
	type args struct {
		m map[string]int64
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumInts(tt.args.m); got != tt.want {
				t.Errorf("SumInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumFloats(t *testing.T) {
	type args struct {
		m map[string]float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumFloats(tt.args.m); got != tt.want {
				t.Errorf("SumFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumIntsOrFloats(t *testing.T) {
	type args struct {
		m map[K]V
	}
	tests := []struct {
		name string
		args args
		want V
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumIntsOrFloats(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumIntsOrFloats() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSumNumbers(t *testing.T) {
	type args struct {
		m map[K]V
	}
	tests := []struct {
		name string
		args args
		want V
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumNumbers(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_main(t *testing.T) {
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
