package findbottomlefttreevalue

import "testing"

func TestFindBottomLeftValue(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test case 1",
			args: args{
				root: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 3,
					},
				},
			},
			want: 1,
		},
		{
			name: "test case 2",
			args: args{
				root: &TreeNode{
					Val: 1,
					Left: &TreeNode{
						Val: 2,
						Left: &TreeNode{
							Val: 4,
						},
						Right: &TreeNode{},
					},
					Right: &TreeNode{
						Val: 3,
						Left: &TreeNode{
							Val: 5,
							Left: &TreeNode{
								Val: 7,
							},
							Right: &TreeNode{},
						},
						Right: &TreeNode{
							Val:   6,
							Left:  &TreeNode{},
							Right: &TreeNode{},
						},
					},
				},
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindBottomLeftValue(tt.args.root); got != tt.want {
				t.Errorf("FindBottomLeftValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
