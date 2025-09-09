package main

// abcabcbb

func lengthOfLongestSubstring(s string) int {
	sLen := len(s)
	visit := make(map[byte]struct{}, 128)
	maxLenght := 0

	left := 0
	right := 0

	for right < sLen {
		currChar := s[right]
		_, has := visit[currChar]

		if !has {
			visit[currChar] = struct{}{}
			if len(visit) > maxLenght {
				maxLenght = len(visit)
			}
			right++
		} else {
			delete(visit, s[left])
			left++
		}
	}

	return maxLenght
}
