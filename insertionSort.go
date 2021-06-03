
package main 

import (
	"fmt"
)

func main (){
	fmt.Printf("Inerstion sort: %d", insertionSort([]int{2,3,-1,-4,9,1,2,0}))
}



// Take j and compare it with all the prevous elements
// We assume that the prevous eleemnts is the biggest elements in the list (ascdening)
func insertionSort(arr[]int)[]int{
	for i:= 1 ; i < len(arr) ; i++{
		j:=i
		for j > 0 && arr[j] < arr[j-1]{
			arr[j], arr[j-1] = arr[j-1], arr[j]
			j--
		}
	}

	return arr
}
