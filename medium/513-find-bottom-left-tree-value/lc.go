package findbottomlefttreevalue

import (
	"container/list"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func FindBottomLeftValue(root *TreeNode) int {
	result := root.Val
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		node := queue.Front().Value.(*TreeNode)
		queue.Remove(queue.Front())

		result = node.Val

		if node.Right != nil {
			queue.PushBack(node.Right)
		}
		if node.Left != nil {
			queue.PushBack(node.Left)
		}
	}

	return result
}
