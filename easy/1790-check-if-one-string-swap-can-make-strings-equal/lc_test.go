package checkifonestringswapcanmakestringsequal

import "testing"

func TestAreAlmostEqual(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test case 1",
			args: args{
				s1: "bank",
				s2: "kanb",
			},
			want: true,
		},
		{
			name: "test case 2",
			args: args{
				s1: "attack",
				s2: "defend",
			},
			want: false,
		},
		{
			name: "test case 3",
			args: args{
				s1: "kelb",
				s2: "kelb",
			},
			want: true,
		},
		{
			name: "test case 4",
			args: args{
				s1: "abcd",
				s2: "dcba",
			},
			want: false,
		},
		{
			name: "test case 5",
			args: args{
				s1: "siyolsdcjthwsiplccjbuceoxmpjgrauocx",
				s2: "siyolsdcjthwsiplccpbuceoxmjjgrauocx",
			},
			want: true,
		}, {
			name: "test case 6",
			args: args{
				s1: "aa",
				s2: "ac",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AreAlmostEqual(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("AreAlmostEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
