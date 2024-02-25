package checkifonestringswapcanmakestringsequal

func AreAlmostEqual(s1 string, s2 string) bool {
	if s1 == s2 {
		return true
	}

	s2b := []byte(s2)

	i := 0

	s2x := -1
	s2y := -1

	for {
		if i == len(s2b) {
			break
		}

		if s1[i] == s2b[i] {
			i++
			continue
		}

		if s2x == -1 {
			s2x = i
			i++
			continue
		}

		s2y = i

		break
	}

	if s2x != -1 && s2y != -1 {
		temp := s2b[s2x]
		s2b[s2x] = s2b[s2y]
		s2b[s2y] = temp
	}

	return s1 == string(s2b)
}
