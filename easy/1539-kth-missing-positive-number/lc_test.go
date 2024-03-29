package kthmissingpositivenumber

import "testing"

func TestFindKthPositive(t *testing.T) {
	type args struct {
		arr []int
		k   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				arr: []int{2, 3, 4, 7, 11},
				k:   5,
			},
			want: 9,
		},
		{
			name: "test case 2",
			args: args{
				arr: []int{1, 2, 3, 4},
				k:   2,
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindKthPositive(tt.args.arr, tt.args.k); got != tt.want {
				t.Errorf("FindKthPositive() = %v, want %v", got, tt.want)
			}
		})
	}
}
