# Answers 

1. Here's the code,
```go
/* ... */
x := []int{48, 96, 86, 68, 57, 82, 63, 70, 37, 34, 83, 27, 19, 97, 9, 17}
smallest := x[0]
for _, num := range x {
  if num < smallest {
    smallest = num
 }
}
fmt.Println("The smallest number in the list is:", smallest)
```

2. These are types of for loop:
- a) `for` loop: This is the most commonly used loop in Go, and is used to iterate over a range of values. It has three parts: initialization, condition, and post iteration.
- b) `range` loop: This loop is used to iterate over elements in an array, slice, map, or channel.
- c) `while` loop: Go doesn't have a separate while loop construct, but you can simulate it using a for loop with only a condition.
- d) `do-while` loop: Like the while loop, Go doesn't have a separate do-while loop, but you can simulate it using a for loop with a break statement.
- e) `infinite` loop: This loop runs indefinitely until it is manually interrupted by the program or terminated by the operating system.

4. Here is the syntax for infinite loop
```go
/* ... */
for {
  fmt.Println("Infinite")
}
```

4. You can terminate a loop in Go using the break statement. When the break statement is encountered inside a loop, the loop immediately terminates and the control flow moves to the statement following the loop.

5. Yes, you can have nested loops in Go. This means that you can have one loop inside another loop. Here's an example of a nested loop in Go,
```go
/* ... */
for i := 0; i < 3; i++ {
    for j := 0; j < 3; j++ {
        fmt.Printf("i: %d, j: %d\n", i, j)
    }
}
```