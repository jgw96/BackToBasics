package main

import (
	"fmt"
	"strconv"
)

func linearSearch(searchValue int) {
	array := [5]int{10, 20, 2, 200, 132}
	for i := 0; i < len(array); i++ {
		if searchValue == array[i] {
			fmt.Println(i)
		}
	}
}

func binarySearch(searchValue int) {
	array := [6]int{2, 10, 20, 132, 200, 205}

	upper := len(array)
	lower := 0
	mid := lower + (upper-lower)/2

	for lower <= upper {
		if array[mid] < searchValue {
			lower = mid + 1
		} else if array[mid] > searchValue {
			upper = mid - 1
		} else if array[mid] == searchValue {
			fmt.Println("value found at" + " " + strconv.Itoa(mid))
			break
		}
	}
}

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

func main() {
	linearSearch(2)
	binarySearch(132)
	insertionSort()
}
