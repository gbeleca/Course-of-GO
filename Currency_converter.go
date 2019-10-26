package main

import (
	"fmt"
)

const dollarRate = 63

func main() {
	//rub перед использованием любую переменную нужно объявить
	var rub float64
	fmt.Println("Конвертер валют из Р в $. Введите интересующую вас сумму рублей в окно, чтобы получить эквивалент в долларах.")
	fmt.Scanln(&rub)
	//если мы хотим присвоить значение переменной без предварительного объявления, нужно использовать конструкцию :=
	result := rub / dollarRate
	fmt.Println("Отдав ", rub, " рублей, вы получите ", result, " $. По курсу ", dollarRate, " рублей за доллар.")
}