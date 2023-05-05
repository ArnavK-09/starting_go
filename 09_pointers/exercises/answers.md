# Answers


1. In Go language, the two pointer operators are `&` and `*`. The `&` operator returns the address of a variable, while the `*` operator is used to access the value that is stored at a particular memory address.

2. To assign a value to a pointer in Go, you can use the `*` operator to access the value at the memory address pointed to by the pointer, and then assign a new value to that variable. For example
```go
/* ... */
var p *int
var i int = 42
p = &i
*p = 21 // i = 21
```

3.
```go
ptrBool := new(bool)
```

4. This program defines a `square` function that takes a `pointer` to a `float32` as its argument. It then updates the value stored at the memory address pointed to by the pointer by squaring it using the `*` operator. In the main function, a `float32` variable `x` is initialized with a value of `6.9`, which is then passed to `square` using the `&` reference operator to take its address. Finally, the updated value of `x` is printed out, which will be **47.61** (the square of `6.9`)

5.
```go
/* ... */
func swapInts(a *int, b *int) {
    tmp := *a
    *a = *b
    *b = tmp
}

func main() {
    var x, y int = 1, 2;
    fmt.Println("Before swap:", x, y)
    swapInts(&x, &y)
    fmt.Println("After swap:", x, y)
}
```

