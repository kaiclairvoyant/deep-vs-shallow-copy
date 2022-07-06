package main

import (
	"fmt"
)

func copySlices() {
	fmt.Println("____________Initiating slice copies____________")
	originalSlice := []int{1}

	fmt.Printf("original slice:\n%v\n", originalSlice)

	shallowCopy := originalSlice
	shallowCopy[0] = 2
	fmt.Printf("copy:\n%v\n", shallowCopy)
	fmt.Printf("original slice:\n%v\n", originalSlice)

	deepCopy := make([]int, len(originalSlice))
	copy(deepCopy, originalSlice)
	deepCopy[0] = 3
	fmt.Printf("copy:\n%v\n", deepCopy)
	fmt.Printf("original slice:\n%v\n", originalSlice)
}
