package runningsumof1darray

func runningSum(nums []int) []int {
	var lastSum int

	for i, v := range nums {
		lastSum += v
		nums[i] = lastSum
	}

	return nums
}
