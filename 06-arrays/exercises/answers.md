# Answers

## Q1
```go
/* ... */
x := array[3] // index starts from 0
```

## Q2
Length = **3**

## Q3

```go
/* ... */
x := [6]string{"a","b","c","d","e","f"}
slice := x[2:]
fmt.Println(slice) // [c d e f]
```

## Q4

You can also use your custom function.

```go
package main

import (
	"fmt"
	"sort"
)

func main() {
	// Define the list of numbers
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}

	// Sort the list in ascending order
	sort.Ints(x)

	// The smallest number is now the first element in the sorted list
	min := x[0]

	// Print the smallest number
	fmt.Println("The smallest number is:", min)
}
```