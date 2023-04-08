package main

/* Go If, else and swtiches */

// some helper pkgs
import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Control Statements in Go")

	// dummy var
	var PogStr string = "Hi"

	// take user input
	var userInput string
	fmt.Print("Enter A Str:- ")
	fmt.Scan(&userInput)
	fmt.Println("You entered, ", userInput)

	fmt.Println("Typeof: ", reflect.TypeOf(userInput).Name())

	fmt.Println() // line space

	/* IF ELSE */
	if userInput == PogStr {
		// if statement
		fmt.Println("POG!")
	} else if reflect.TypeOf(userInput).Name() == "string" {
		// else if statement
		fmt.Println("Gotta String")
	} else {
		// else statement
		fmt.Println("Got Ntg")
	}

	/* Same thing with SWITCH */
	switch userInput {
	case PogStr:
		{
			fmt.Println("Pog")
		}
	default:
		{
			// if mone of the statements match
			fmt.Println("Got Smth In Str")
		}
	}

}
