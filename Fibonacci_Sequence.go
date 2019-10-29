package main

import (
	"fmt"
)

func main() {
	fmt.Println("Эта функция покажет 100 первых чисел Фибоначи:")
	var Fibonacci_Sequence []int
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			Fibonacci_Sequence[i] = i + j
		}
	}

}
