package concatenationofarray

func GetConcatenation(nums []int) []int {
	ans := make([]int, 2*len(nums))

	copy(ans, nums)
	copy(ans[len(nums):], nums)

	return ans
}
