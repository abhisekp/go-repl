package bracketmatching

import (
	"fmt"
	"github.com/abhisekp/go-repl/stackusingqueue"
	"strings"
)

func Run() {
	fmt.Println("()[]{}:", MatchBrackets("()[]{}")) // true
	fmt.Println("({[}]):", MatchBrackets("({[}])")) // false
	fmt.Println("([{}]):", MatchBrackets("([{}])")) // true
	fmt.Println("}][{", MatchBrackets("}][{"))      // false
}

func MatchBrackets(str string) bool {
	// is Odd
	if len(str)%2 != 0 {
		return false
	}

	matchMap := map[rune]rune{
		'}': '{',
		']': '[',
		')': '(',
	}

	bracketsStack := stackusingqueue.NewStack[rune]()

	for _, ch := range str {
		if strings.ContainsRune("{[(", ch) {
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
