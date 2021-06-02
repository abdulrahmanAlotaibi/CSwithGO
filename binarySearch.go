package main

import (
	"fmt"
	"math"
)


func main (){
	// The list needs to be sorted
	fmt.Println(binarySearch([]int{2,4,8,9}, 9))
}

func binarySearch(list []int, el int) (int) {
	
	low := 0
	high := len(list) - 1


	for low <= high {

		mid := int(math.Floor(float64((low + high)) / 2.0))

		guess := list[mid]

		if guess == el {
			return mid
		} else if guess < el {
			low = mid + 1
		} else  {
			high = mid - 1 
		}
	}

	return -1
}
