package sort

import (
	"fmt"
	"time"
)

func Obubble(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arrCopy[j] > arrCopy[j+1] {
				arrCopy[j], arrCopy[j+1] = arrCopy[j+1], arrCopy[j]
			}
		}
	}
	end := time.Now().UnixNano()
	// fmt.Println(arrCopy)
	fmt.Println("Obubble  takes:  ", end-start)
	return end - start
}

func Abubble2(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()
	for i := 1; i < len(arrCopy); i++ {
		for j := 0; j < len(arrCopy)-i; j++ {
			if arrCopy[j] > arrCopy[j+1] {
				arrCopy[j], arrCopy[j+1] = arrCopy[j+1], arrCopy[j]
			}
		}
	}
	end := time.Now().UnixNano()
	// fmt.Println(arrCopy)
	fmt.Println("Abubble2  takes:  ", end-start)
	return end - start
}

func Abubble3(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()
	swaped := true
	n := len(arrCopy)
	for swaped {
		swaped = false
		for j := 0; j < n-1; j++ {
			if arrCopy[j] > arrCopy[j+1] {
				arrCopy[j], arrCopy[j+1] = arrCopy[j+1], arrCopy[j]
				swaped = true
			}
		}
		n--
	}
	end := time.Now().UnixNano()
	// fmt.Println(arrCopy)
	fmt.Println("Abubble3  takes:  ", end-start)
	return end - start
}

func Abubble4(arr []int) int64 {
	arrCopy := make([]int, len(arr), len(arr))
	copy(arrCopy, arr)
	start := time.Now().UnixNano()
	n := len(arrCopy) - 1
	for n != 0 {
		newN := 0
		for j := 0; j < n; j++ {
			if arrCopy[j] > arrCopy[j+1] {
				arrCopy[j], arrCopy[j+1] = arrCopy[j+1], arrCopy[j]
				newN = j
			}
		}

		n = newN
	}

	end := time.Now().UnixNano()
	// fmt.Println(arrCopy)
	fmt.Println("Abubble4  takes:  ", end-start)
	return end - start
}
