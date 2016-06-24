package main

import "fmt"

var min int
var tmp int

func insertionSort() {
	array := [6]int{4, 5, 2, 200, 19, 45}

	/*
	 * Whenever the value in the sorted section is greater than the value
	 * in the unsorted section, shift all items in the sorted section over
	 * by one. This creates space in which to insert the value.
	 */
	for i := 0; i < len(array); i++ {
		min = i
		for j := i + 1; j < len(array); j++ {
			if array[j] < array[min] {
				min = j
			}
		}
		tmp = array[i]
		array[i] = array[min]
		array[min] = tmp
	}

	fmt.Println(array)
}