package main

// imports
import "fmt"

func main() {
	fmt.Println("\nThe only difference between this and an array is the missing length between the brackets. In this case x has been created with a length of 0.\n")

	/* len(slice) < len(array) */

	// simple slice
	var x []int
	// with inbuilt function with length 5
	y := make([]int, 5)

	fmt.Println(x, y)

	// another way
	simpleArray := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 0,
	}

	z := simpleArray[0:5] // [low:high]
	a := simpleArray[3:]  // high = len(simpleArray)
	fmt.Println(z, a)

	// some functions

	// append something
	fmt.Println(append(z, 6), z) //z remain unchanged

	// copy
	a1 := []int{
		1, 2, 3,
	}
	a2 := make([]int, 4)
	a2 = append(a2, 11) // more data
	copy(a2, a1)        // copy a1 contents to a2
	fmt.Println(a1, a2) // array copied, but only for data present

}
