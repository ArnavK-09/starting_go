# Answers

1.
```go
package main

import "fmt"

func main() {
	result := addAll([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Print(result)
}

func addAll(list []int) (sum int) {
	for _, num := range list {
		sum += num
	}
	return
}
```

2.
```go
/* ... */
func half(num int) (halfValue int, status bool) {
	// halved value
	halfValue = num / 2
	// check
	if halfValue%2 == 0 {
		status = true
	} else {
		status = false
	}
	// return
	return
}
```

3.
```go
/* ... */
func findMax(nums ...int) (max int) {
	max = nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return
}
```


