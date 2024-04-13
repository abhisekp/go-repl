package bracketmatching

import (
	"github.com/abhisekp/go-repl/stackusingqueue"
	"slices"
)

func Run() {

}

func MatchBrackets(str string) bool {
	if len(str)%2 != 0 {
		return false
	}

	matchMap := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	bracketsStack := stackusingqueue.NewStack[rune]()

	startingBrackets := []rune{'{', '[', '('}

	for _, ch := range str {
		if slices.Contains(startingBrackets, ch) {
			bracketsStack.Push(ch)
			continue
		}

		if startBracket, ok := matchMap[ch]; ok {
			if bracketsStack.Peek() != startBracket {
				return false
			}
			bracketsStack.Pop()
		}
	}

	return bracketsStack.IsEmpty()
}
