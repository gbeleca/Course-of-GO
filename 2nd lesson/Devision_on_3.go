package main

import (
	"fmt"
)

func main(userNumber) {
	for i := 0; i <= 10; i++ {
		fmt.Println("Введите число, чтобы узнать является ли оно четным:")
		var userNumber int
		fmt.Scanln(&userNumber)
		var residue int = userNumber % 3
		if userNumber%3 == 0 {
			fmt.Println("Это число делится без остатка на 3.")
		} else {
			fmt.Println("Это число не делится на 3, остаток:", residue)
		}
	}

}
