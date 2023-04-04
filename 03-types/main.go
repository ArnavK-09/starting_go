package main

// imports
import (
	"fmt"
)

/**
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128
*/

func main() {
	fmt.Println("Types in golang")

	// boolean (true/false)
	var boolean bool // init with false
	fmt.Println(boolean)

	// int
	var (
		a int      // init with 0
		b int = -1 //siigned int
		// Err: c int = -1
		c uint       = 1 //uint cant have - value
		d complex128 = 0.867 + 0.5i
		e rune
	)

	fmt.Println(a, b, c, d)

	// float
	const floating float64 = 69.8
	fmt.Println(floating)

	// strings
	var str string //init with ""
	fmt.Print(str)

	// type conversion
	// type(to be converted)

	fmt.Printf("e is of type %T\n", e)
	var converted string = string(e)
	fmt.Printf("e (converted) is of type %T\n", converted)

	fmt.Printf("e is of type %T\n", e) // main variable remain unchanged
}
