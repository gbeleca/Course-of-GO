package main

import (
	"fmt"
)

func main() {
	fmt.Println("Эта функция покажет 100 первых чисел Фибоначи:")
	var Fibonacci_Sequence [] int
	var Fibonacci_Sequence[0]=0
	for i := 0; i < 100; i++ {
		for j := 1; j < 100; j++ {
			Fibonacci_Sequence[j] = i + j
		}
	}

}
