package main

import (
	"fmt"
	"strconv"
)

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
