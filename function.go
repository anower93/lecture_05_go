package main

import "fmt"

// Example 1

/*
func add(x int, y int) int{
	result := x + y

	return result
}
*/

// Example 2, We can write variable type only one time if it all are same following this example with parameter
/*
func add(x, y int) int {
	result := x + y

	return result
}
*/

// Example 3:

/*

func add(x, y int) (result int) {
	result = x + y

	return result
}
*/

//Example 4

/*
func add(x, y int) (result int) {

	result = x + y

	return
}

*/

// Example 5:

// func doublecalculation(x, y int) (p, a int) {

// 	p = 2 * (x + y)
// 	a = x + y
// 	return
// }

func main() {
	fmt.Println("Welcome to Function!")

	// This is Anonymous function
	a := func(x, y int) (r int) {
		r = x * y
		return
	}(10, 10)

	fmt.Println(a)

}
