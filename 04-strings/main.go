package main

/* Immutable Strings */

// imports
import (
	"fmt"
	"strings"
)

func main() {
	// ways to declare strings
	var a string = "String Value"
	const b string = "Another String"
	c := "Yeah String"
	const (
		d string = "TOO STRING"
		e        = "Hmmm"
	)
	var (
		f string = "Too String"
		g        = "Hmmm"
	)
	var h string      // empty string
	i := []byte("Hi") // converts steing to array of bytes

	fmt.Println(a, b, c, d, e, f, g, h, i)

	// Err: a[0] = "P"

	i[0] = 102                // 102 = f
	fmt.Println(i, string(i)) // get value of i bytes

	fmt.Println(b[3]) // get any letter at index of string

	// loop string
	for i, x := range "STRING!" {
		fmt.Println(x, " ", i)
	}

	//length of string
	fmt.Println("Length of string f is: ", len(f))

	/* string packages small functions*/
	fmt.Println("ToLower ", strings.ToLower(d))
	fmt.Println("ToUpper ", strings.ToUpper(d))
	fmt.Println("ReplaceAll ", strings.ReplaceAll(e, "m", "uh"))
	fmt.Println("Repeat ", strings.Repeat("Hi! ", 3))
	fmt.Println("Split ", strings.Split(c, " "))
	fmt.Println("Index ", strings.Index(c, "Ye")) // -1 = not exist
	fmt.Println("Contains ", strings.Contains("b", "x"))

	/* All Methods: https://pkg.go.dev/strings */

}

/*

\\	Backslash(\)

\000	Unicode character with the given 3-digit 8-bit octal code point

\’	Single quote (‘). It is only allowed inside character literals

\”	Double quote (“). It is only allowed inside interpreted string literals

\a	ASCII bell (BEL)

\b	ASCII backspace (BS)

\f	ASCII formfeed (FF)

\n	ASCII linefeed (LF)

\r	ASCII carriage return (CR)

\t	ASCII tab (TAB)

\uhhhh	Unicode character with the given 4-digit 16-bit hex code point.
 	Unicode character with the given 8-digit 32-bit hex code point.

\v	ASCII vertical tab (VT)

\xhh	Unicode character with the given 2-digit 8-bit hex code point.

*/
