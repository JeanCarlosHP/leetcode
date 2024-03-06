package consecutivecharacters

import "testing"

func Test_maxPower(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				s: "leetcode",
			},
			want: 2,
		},
		{
			name: "test case 2",
			args: args{
				s: "abbcccddddeeeeedcba",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPower(tt.args.s); got != tt.want {
				t.Errorf("maxPower() = %v, want %v", got, tt.want)
			}
		})
	}
}
