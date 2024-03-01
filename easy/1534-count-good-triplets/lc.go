package countgoodtriplets

func CountGoodTriplets(arr []int, a int, b int, c int) (gt int) {
	aLen := len(arr)

	for i := 0; i < aLen; i++ {
		for j := i + 1; j < aLen; j++ {
			if int(abs(arr[i]-arr[j])) <= a {
				for k := j + 1; k < aLen; k++ {
					c2 := int(abs(arr[j]-arr[k])) <= b && int(abs(arr[i]-arr[k])) <= c
					if c2 {
						gt++
					}
				}

			}
		}
	}

	return gt
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
