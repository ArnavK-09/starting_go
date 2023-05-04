package main

// imports
import "fmt"

// main function
func main() {
	fmt.Println("Functions in go\n")
	result := idk([]interface{}{5, true, "hi"})
	fmt.Println(result)

	// demo variables
	var num1, num2 int = 5, 10
	fmt.Println(num1, num2)

	// swaping values
	swap(&num1, &num2)
	fmt.Println(num1, num2)

	// nested
	fmt.Println(f1())

	// multiple
	fmt.Println(yeah())
	val1, val2 := yeah()
	fmt.Println("Values: ", val1, val2)

	// Variadic
	showArgs("Hii", 69, 3.3, true)

	// closure
	fmt.Println(guessFunc())

	// function in variable
	guessAgain := func() {
		fmt.Println("Guessed")
	}
	guessAgain()

	// another closure
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) // 0
	fmt.Println(nextEven()) // 2

	// recursion
	run(1)

	// defer : schedules function to last
	defer first()
	second()

	// error handler
	defer func() {
		whatsErr := recover()
		fmt.Println("Error captured:-\n", whatsErr)
	}()
	panic("Errror Intentionally")

}

// custom functions
func idk(arg []interface{}) (a string) {
	a = "test" // named return
	return
	// panic("IDK") : custom err in function
}

// values with pointerz
func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

// nested functions
func f1() string {
	return f2()
}

func f2() string {
	return "hiii"
}

// returning multiple values
func yeah() (bool, bool) {
	return true, false
}

// Variadic functions
func showArgs(args ...interface{}) {
	fmt.Println(args)
}

// closure
func guessFunc() func() uint {
	return func() uint {
		return 69
	}
}

// another closure
func makeEvenGenerator() func() uint {
	i := uint(0)

	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

// demo functions for defer use
func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}

// recursion
func run(n uint) {
	fmt.Println("Value of n: ", n)
  // inc
	n++
	if n != 5 {
		run(n)
	} else {
		fmt.Println("Last recured value: ", n)
	}
}
