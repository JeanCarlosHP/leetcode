package concatenationofarray

import (
	"reflect"
	"testing"
)

func TestGetConcatenation(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 2, 1},
			},
			want: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name: "test case 1",
			args: args{
				nums: []int{1, 3, 2, 1},
			},
			want: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetConcatenation(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConcatenation() = %v, want %v", got, tt.want)
			}
		})
	}
}
