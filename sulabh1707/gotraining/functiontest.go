package main

import "fmt"

//Here’s a function that takes two ints and returns their sum as an int.
func plus(a int, b int) int {
	//Go requires explicit returns, i.e. it won’t automatically return the value of the last expression.
	return a + b
}

//When you have multiple consecutive parameters of the same type, you may omit the type name for the like-typed parameters up to the final parameter that declares the type.
func plusPlus(a, b, c int) int {
	return a + b + c
}

func getSum(arr []int, size int) int {
	var i, sum int

	for i = 0; i < size; i++ {
		sum += arr[i]
	}
	return sum
}

func main() {
	//Call a function just as you’d expect, with name(args).
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
	res = plusPlus(1, 2, 3)
	fmt.Println("1+2+3 =", res)

	var array = []int{1, 2, 3, 4, 5}
	size := len(array)
	fmt.Print("size:", size)
	sumarry = getSum(array:=[]int{1,2,3,4,5,6}, size:=len(array))
	fmt.Println(sumarry)
}
