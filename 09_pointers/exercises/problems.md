# Problems

1. What are two pointer operators?

2. How do you assign a value to a pointer?

3. How do you create a new pointer with type of `bool` with `new` function?

4. What is the value of `x` after running this program:
```go
/* ... */
func square(x *float32) {
 *x = *x * *x
}
func main() {
 var x float32 = 6.9;
 square(&x)
}
```

5. Write a program that can swap two integers value `(x := 1; y := 2; swap(&x, &y)` should give you `(x=2 and y=1)`
