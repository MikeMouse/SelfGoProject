package sort

import (
	"fmt"
	"time"
)

func CocktailSort(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()
	left, right := 0, len(arrCopy)-1
	// swaped := false
	for left < right {
		for i := left; i < right; i++ {
			if arrCopy[i] > arrCopy[i+1] {
				arrCopy[i], arrCopy[i+1] = arrCopy[i+1], arrCopy[i]
				// swaped = true
			}
		}

		right--
		for i := right; i > 0; i-- {
			if arrCopy[i-1] > arrCopy[i] {
				arrCopy[i-1], arrCopy[i] = arrCopy[i], arrCopy[i-1]
			}
		}
		left++
	}

	end := time.Now().UnixNano()
	// fmt.Println(arrCopy)
	fmt.Println("Abubble4  takes:  ", end-start)
	return end - start
}
