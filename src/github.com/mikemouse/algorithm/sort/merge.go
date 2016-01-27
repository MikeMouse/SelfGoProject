package sort

import (
	"fmt"
	"time"
)

func MergeSort(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()

	fmt.Println(arrCopy)
	result := make([]int, len(arrCopy), len(arrCopy))
	merge(arrCopy, result, 0, len(arrCopy)-1)

	end := time.Now().UnixNano()
	fmt.Println(arrCopy)
	fmt.Println("MergeSort  takes:  ", end-start)
	return end - start
}

func merge(arr []int, result []int, start int, end int) {
	fmt.Println("merge :", start, "/", end)
	if start >= end {
		return
	}

	mid := (end-start)/2 + start
	leftStart := start
	leftEnd := mid
	rightStart := mid + 1
	rightEnd := end
	merge(arr, result, leftStart, leftEnd)
	merge(arr, result, rightStart, rightEnd)
	left := leftStart
	right := rightStart
	index := leftStart
	for left <= leftEnd && right <= rightEnd {
		if arr[left] > arr[right] {
			result[index] = arr[right]
			right++
		} else {
			result[index] = arr[left]
			left++
		}
		index++
	}

	// fmt.Println("merge1:", result)
	for left <= leftEnd {
		result[index] = arr[left]
		index++
		left++
	}
	// fmt.Println("merge2:", result)

	for right <= rightEnd {
		result[index] = arr[right]
		index++
		right++
	}
	// fmt.Println("merge3:", result)

	for i := start; i <= end; i++ {
		arr[i] = result[i]
	}

	fmt.Println("len more than 2  end:", arr)

}
