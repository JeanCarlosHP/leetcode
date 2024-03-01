package kthmissingpositivenumber

func FindKthPositive(arr []int, k int) int {
	var missingInts = []int{}
	var i int

	for i := 0; i < k; i++ {
		arr = append(arr, 0)
	}

	for {
		if i+1 != arr[i] {
			missingInts = append(missingInts, i+1)
			arr = append(arr, 0)
			copy(arr[i+1:], arr[i:])
		}

		if len(missingInts) == k {
			break
		}

		i++
	}

	return missingInts[len(missingInts)-1]
}
