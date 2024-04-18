package canbepalindrome

func Run() {
	CanBePalindrome("amadm")
}

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

	return oddCount <= 1
}
