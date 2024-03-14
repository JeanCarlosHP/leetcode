package stringmatchinginanarray

import "strings"

// Tests will be not pass, because order of list is random
func stringMatching(words []string) []string {
	if len(words) < 2 {
		return []string{}
	}

	m := make(map[string]struct{})

	for _, word1 := range words {
		for _, word2 := range words {
			if word1 != word2 {
				if strings.Contains(word2, word1) {
					m[word1] = struct{}{}
				}
			}
		}
	}

	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}
