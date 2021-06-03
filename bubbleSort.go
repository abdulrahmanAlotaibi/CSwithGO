package main

import (
	"fmt"
)

func main(){
	fmt.Println(bubbleSort([]int{5,2,3,22,-4,0,11,-5,1,1}))
	
}

func bubbleSort(arr []int )[]int{
	for i:= 0 ; i < len(arr) - 1 ; i++{
		for j:=0 ; j < len(arr) - 1 - i ; j++ {
			if arr[j] > arr[j + 1]{
				arr[j], arr[j+1] = arr[j+1],arr[j]
			} 		
		} 
	}
	return arr
}mv
