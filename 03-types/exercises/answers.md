# Answers

## Q1
```go
/* ... */
const result rune = 12345 * 54321
fmt.Println(result)
```

## Q2
```go
/* ... */
str := "Hello, world!"
length := len(str)
fmt.Println(length)
```

## Q3
The first two expressions evaluate to `false`, but the third expression `!(false && false)` evaluates to `true` because `false && false` is `false` and `!false` is `true`. Therefore, the overall expression becomes `false || false || true`, which is **true**.

## Q4
The default value of a bool variable in Go is ` false `

## Q5
In Go, an int64 is a signed integer type that can represent whole numbers between -(2^63) and 2^63-1. The maximum value that can be stored in an int64 variable in Go is 2^63-1, which is equal to 9223372036854775807.