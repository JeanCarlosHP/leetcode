package maketwoarraysequalbyreversingsubarrays

func canBeEqual(target []int, arr []int) bool {
	m := make(map[int]int)

	for _, v := range target {
		m[v]++
	}

	for _, v := range arr {
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			return false
		}
	}

	return true
}
