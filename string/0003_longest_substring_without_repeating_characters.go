package string

func lengthOfLongestSubstring(s string) int {
	letters := make(map[byte]int)
	checkPoint := 0
	length := 0

	for i := range s {
		pos, ok := letters[s[i]]
		if ok {
			if length < i-checkPoint {
				length = i - checkPoint
			}
			if checkPoint < pos+1 {
				checkPoint = pos + 1
			}
		}
		letters[s[i]] = i
	}

	if length < len(s)-checkPoint {
		length = len(s) - checkPoint
	}

	return length
}
