package sort

import (
	"fmt"
	"time"
)

func InsertSort(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()

	for i := 1; i < len(arrCopy); i++ {
		temp := arrCopy[i]
		j := i - 1
		for ; j >= 0 && arrCopy[j] > temp; j-- {
			arrCopy[j+1] = arrCopy[j]
		}
		arrCopy[j+1] = temp
	}
	end := time.Now().UnixNano()
	fmt.Println(arrCopy)
	fmt.Println("InsertSort  takes:  ", end-start)
	return end - start
}
