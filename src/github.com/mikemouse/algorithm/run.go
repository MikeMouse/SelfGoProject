package main

import (
	"fmt"
	"github.com/mikemouse/algorithm/sort"
	"math/rand"
)

// init array
// var arr = randArray(10000)

var arr = []int{9, 2, 7, 4, 6, 5, 8, 1, 3}

func main() {
	sort.Obubble(arr)
	sort.Abubble2(arr)
	sort.Abubble3(arr)
	sort.Abubble4(arr)

	fmt.Println("==================")
	sort.CocktailSort(arr)
	fmt.Println("==================")
	sort.InsertSort(arr)
	fmt.Println("==================")
	sort.MergeSort(arr)
	fmt.Println("==================")
}

func randArray(count int) []int {
	if count <= 0 {
		return nil
	}

	result := make([]int, count, count)
	// start := time.Now().UnixNano()
	for i := 0; i < count; i++ {
		result[i] = rand.Intn(count)
	}
	// end := time.Now().UnixNano()
	// fmt.Println("Random array initial end at : ", end)
	// fmt.Println("Random array initial takes :  %v nano", end-start)
	return result
}
