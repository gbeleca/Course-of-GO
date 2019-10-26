package main

import "fmt"

func main() {
	fmt.Println("Конвертер валют из Р в $. Введите интересующую вас сумму рублей в окно, чтобы получить эквивалент в долларах.")
	var rur float64
	fmt.Scanln(&rur)

	const dollarRate float64 = 63.68

	var result float64 = rur / dollarRate

	fmt.Printf("Отдав %v рублей, вы получите %.2f $. По курсу %v рублей за доллар.", rur, result, dollarRate)
}
