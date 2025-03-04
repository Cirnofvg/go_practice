package fibonacci

import "fmt"

func Fibonacci(n int) {
	i, j := 0, 1
	for k := 0; k < n; k = k + 1 {
		i, j = i+j, i
		fmt.Println(j)
	}
}
