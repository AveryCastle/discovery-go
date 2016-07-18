package main

import "fmt"

func fibonachi(n int) int {
	p, q := 0, 1
	for i := 0; i < n; i++ {
		fmt.Println(p)
		p, q = q, p+q
	}
	return p
}

func main() {
	result := fibonachi(10)
	fmt.Println(result)
}
