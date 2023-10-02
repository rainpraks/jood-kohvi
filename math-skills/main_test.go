package main

import "testing"

func TestMedian(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		// TODO: Add test cases.
		{"test1", args{[]int{1, 2, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Median(tt.args.data); got != tt.want {
				t.Errorf("Median() = %v, want %v", got, tt.want)
			}
		})
	}
}
