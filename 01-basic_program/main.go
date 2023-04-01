package main
// Every go file should be started with package declaration, files with same package bundles

// import packages
import "fmt"  // in built modules
import "math"
/* Or import together
   import (
     "fmt";
     "math"
   )
*/

// main function
func main() {
  // basic print statement in go 
  fmt.Println("Using imported packages for our work\n", "Value Of Pi: ", math.Pi)
}