package minimum_subarray

import "testing"

func TestMinimumSubArray(t *testing.T) {
	type args struct {
		inputs []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "TC1", args: args{inputs: []int{2, 3, 1, 2, 4, 3}, target: 7}, want: 2},
		{name: "TC1", args: args{inputs: []int{1, 4, 4}, target: 4}, want: 1},
		{name: "Zero", args: args{inputs: []int{2, 3, 1, 2, 4, 3}, target: 22}, want: 0},
		{name: "One", args: args{inputs: []int{2, 3, 1, 2, 4, 3}, target: 1}, want: 1},
		{name: "Single", args: args{inputs: []int{2}, target: 1}, want: 1},
		{name: "Failed", args: args{inputs: []int{1, 2, 3, 4, 5}, target: 15}, want: 5},
		{name: "Failed2", args: args{inputs: []int{2, 3, 1, 1, 1, 1, 1}, target: 5}, want: 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumSubArray(tt.args.inputs, tt.args.target); got != tt.want {
				t.Errorf("MinimumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
