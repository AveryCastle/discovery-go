package main

import "fmt"

func Move(from int, to int, via int, n int) {
	if n < 1 {
		return
	}

	Move(from, via, to, n-1)
	fmt.Println(from, " -> ", to)
	Move(via, to, from, n-1)
}

func Hanoi(n int) {
	Move(1, 2, 3, 3)
}

func main() {
	Hanoi(3)
}
