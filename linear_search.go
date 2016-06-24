package main

import (
	"fmt"
)

func linearSearch(searchValue int) {
	array := [5]int{10, 20, 2, 200, 132}
	for i := 0; i < len(array); i++ {
		if searchValue == array[i] {
			fmt.Println(i)
		}
	}
}
