# Answers 

## Q1 

```bash
Small
```

## Q2 

```go
i := 10
switch {
case i > 10:
	fmt.Println("Big")
default:
	fmt.Println("Small")
}
```

## Q3

Yes, you can have nested if-else 
```go
/* ... */
x := 10
y := 20

if x > 0 {
  fmt.Println("x is positive")
  if y > 0 {
    fmt.Println("y is also positive")
  } else {
    fmt.Println("y is not positive")
  }
} else {
  fmt.Println("x is not positive")
}
```