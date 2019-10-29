package main

import (
	"fmt"
)

func main(userNumber) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Введите число, чтобы узнать является ли оно четным:")
		var userNumber int
		fmt.Scanln(&userNumber)
		if userNumber%2 == 0 {
			fmt.Println("Это число четное.")
		} else {
			fmt.Println("Это число четное, или не является числом.")
		}
	}

}
