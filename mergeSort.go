package main

import (
	"fmt"
)

func main(){
	fmt.Println(mergeSort([]int{2,1,-3,0}))
}


func mergeSort(arr []int)[]int{
	
	n := len(arr)

	if n == 1 {
		return arr
	}
	
	mid := int(n / 2)
	
	right := arr[0: mid]
	
	left := arr[mid : n]
	fmt.Println(right, left)
	return merge(mergeSort(right), mergeSort(left))
	
}



func merge(right[]int, left[]int) []int{
	sorted := make([]int, len(right) + len(left))
	
	i := 0
	
	for len(right) > 0  && len(left) > 0 {
		if right[0] < left[0]{
			sorted[i] = right[0]
			right = right[1:]
		}else {
			sorted[i] = left[0]
			left=left[1:]
		}
		i++
	}

	for j:= 0; j < len(right) ; j++ {
		sorted[i] =  right[j]
		i++
	}

	for j:= 0 ; j < len(left) ; j++ {
		sorted[i] = left[j]
		i++
	}

	fmt.Println(sorted)
	
	return sorted
}



