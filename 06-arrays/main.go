package main

// imports
import "fmt"

func main() {
	// simple array with strings and limit of 5
	var array [5]string
	fmt.Println(array) // init blank array eith default value of string ie ""

	// array length
	fmt.Println(len(array)) // 5

	//Err: array[5] = "Hey" || index starts from 0
	array[4] = "Hey"
	fmt.Println(array)

	// another method
	arr := [4]int{
		1,
		2,
		3, // comma
		// 4,
	}
	fmt.Println(arr)

	// array with any item
	array2 := [5]interface{}{
		1,
		true,
		"Hii",
		69.5,
		arr, // even arrays
	}
	fmt.Println(array2)

	//nested arrays
	arr2 := [][4]int{
		arr,
		arr,
		arr,
		arr,
		[4]int{ // declaring new arrays in array
			len(arr),
		},
  }
	fmt.Println(arr2)
}
