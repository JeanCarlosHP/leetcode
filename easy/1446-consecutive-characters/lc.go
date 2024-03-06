package consecutivecharacters

func maxPower(s string) int {
	var lastRune rune
	var i = -1
	arr := make([]int, len(s))

	for _, w := range s {
		if w == lastRune {
			arr[i]++
			continue
		}

		lastRune = w
		i++
		arr[i]++
	}

	var b int

	for _, v := range arr {
		if v > b {
			b = v
		}
	}

	return b
}
