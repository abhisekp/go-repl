package histogramarea

import (
	"fmt"
	"github.com/abhisekp/go-repl/stack"
)

func Run() {
	buildingHeights := []Size{{6, 1}, {2, 1}}
	fmt.Println(findMaxArea(buildingHeights)) // 6

	buildingHeights = []Size{
		{6, 1},
		{2, 1},
		{5, 1},
		{4, 1},
		{5, 1},
		{1, 1},
		{6, 1},
	}
	fmt.Println(findMaxArea(buildingHeights)) // 12

	buildingHeights = []Size{
		{6, 1},
		{2, 1},
		{5, 1},
		{4, 1},
		{6, 1},
	}
	fmt.Println(findMaxArea(buildingHeights)) //

	buildingHeights = []Size{
		{6, 1},
		{2, 1},
		{5, 1},
		{2, 1},
		{6, 1},
	}
	fmt.Println(findMaxArea(buildingHeights)) // 14
}

type Size struct {
	Height int
	Width  int
}

func findMaxArea(heights []Size) int {
	maxArea := 0
	width := 1
	minIdxStack := stack.NewStack[int]()

	for i, size := range heights {
		width += size.Width
		area := size.Height * size.Width
		maxArea = max(maxArea, area)
		if minIdxStack.IsEmpty() || size.Height >= heights[minIdxStack.Peek()].Height {
			minIdxStack.Push(i)
		} else {
			top := minIdxStack.Peek()
			minIdxStack.Pop()
			topSize := heights[top]
			if minIdxStack.IsEmpty() {

			}
		}
	}

	return maxArea
}
