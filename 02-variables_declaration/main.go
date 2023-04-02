package main

// go inbuilt pkgs
import "fmt"
import "reflect"

func main() {
	fmt.Println("Variables Declaration In Go")

	// basic declaration with data type
	// var (name) (type)
	var num int
	// one liner
	var ok bool = true

	fmt.Println(num, ok) // int init with 0

	// setting values
	num = 69
	fmt.Println(num)

	// const - immutable declaration
	const str string = "Hellow"
	// Err: str = "change it"
	fmt.Println(str)

	// declaring multiple variables ( also works with 'const' )
	var (
		a int    = 68
		b string = "yes"
		c bool   // with default value
		d = 77.4 // automatic picked data type
	)

	// Err: d = "try" (cant change data type)
	fmt.Println(a, b, c, d)

	// Declaring/Init variable with automatic data type
	isOk := true
	// Err: isOk = false
	fmt.Println(isOk)

	// tip
	log := fmt.Println
	log("easy print")

	// finding type of object
	log(reflect.TypeOf(a))
	log(reflect.TypeOf(d))

}
