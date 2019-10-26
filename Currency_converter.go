package main

import (
	"fmt"
)

const dollarRate = 63e68

func getResult() {
	fmt.Println("Конвертер валют из Р в $. Введите интересующую вас сумму рублей в окно, чтобы получить эквивалент в долларах.")
	fmt.Scanln(&rub)
	result = rub / dollarRate
	fmt.Println("Отдав ", rub, " рублей, вы получите ", result, " $. По курсу ", dollarRate, " рублей за доллар.")
}
