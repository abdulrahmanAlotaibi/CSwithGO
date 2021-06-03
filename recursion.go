package main 

import "fmt"

func main (){
	// fibonacci sequence
	// See https://en.wikipedia.org/wiki/Fibonacci_number
	fmt.Println(fibonacci(5))

	// Sum all the numbers between index 0 to n-1
	fmt.Println(binarySum([]int {1,10,3,5,9,-2}))


	// Finding Factorial number 4!
	fmt.Println(factorial(4))

}

func binarySum(arr[]int)int{

	if len(arr) == 1 {
		return arr[0]
	}

	mid := int( len(arr) / 2 )

	return binarySum(arr[:mid]) + binarySum(arr[mid:len(arr)])
}

func fibonacci(n int)int{

	if n < 2 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func factorial(n int)int{
	if n==1 {
		return 1
	}
	return n * factorial(n-1)
}
