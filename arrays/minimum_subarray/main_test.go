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
		{name: "Zero", args: args{inputs: []int{2, 3, 1, 2, 4, 3}, target: 8}, want: 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MinimumSubArray(tt.args.inputs, tt.args.target); got != tt.want {
				t.Errorf("MinimumSubArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
