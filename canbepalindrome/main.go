package canbepalindrome

func CanBePalindrome(str string) bool {
	freq := map[rune]int{}

	for _, ch := range str {
		freq[ch]++
	}

	var oddCount int
	for _, fr := range freq {
		if fr%2 != 0 {
			oddCount++
		}
	}

	if len(str)%2 == 0 {
		return oddCount == 0
	} else {
		return oddCount%2 != 0
	}
}
