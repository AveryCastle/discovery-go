package main

import "fmt"

func printSong(amount int) {
	fmt.Printf("타잔이 %d원짜리 팬티를 입고, %d원짜리 칼을 차고 노래를 한다.\n", amount, (amount + 10))
}

func calcTazanSong(n int) int {
	amount := 0

	for n > 0 {
		amount += 10
		n--
		printSong(amount)
	}

	return n
}

func main() {
	calcTazanSong(5)
}
