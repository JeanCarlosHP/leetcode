package maximumunitsonatruck

import "testing"

func TestMaximumUnits(t *testing.T) {
	type args struct {
		boxTypes  [][]int
		truckSize int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				boxTypes:  [][]int{{1, 3}, {2, 2}, {3, 1}},
				truckSize: 4,
			},
			want: 8,
		},
		{
			name: "test case 2",
			args: args{
				boxTypes:  [][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}},
				truckSize: 10,
			},
			want: 91,
		},
		{
			name: "test case 3",
			args: args{
				boxTypes:  [][]int{{4, 2}, {5, 5}, {4, 1}, {1, 4}, {2, 5}, {1, 3}, {5, 3}, {1, 5}, {5, 5}, {1, 1}},
				truckSize: 24,
			},
			want: 95,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaximumUnits(tt.args.boxTypes, tt.args.truckSize); got != tt.want {
				t.Errorf("MaximumUnits() = %v, want %v", got, tt.want)
			}
		})
	}
}
