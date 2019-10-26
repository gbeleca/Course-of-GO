package TriangleCulc

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Вычисление площади, периметра и гипотенузы по катетам прямоугольного треугольника. Введите катеты:")
	fmt.Scanln(&AB)
	fmt.Scanln(&AC)
	var BC float64 = math.Sqrt(math.Pow(AB, 2) + math.Pow(AC, 2))
	var P float64 = AB + AC + BC
	var S float64 = (AB * AC) / 2
	var H float64 = math.Sqrt(math.Pow(AB, 2) - math.Pow(AC, 2))

	fmt.Printf("При катетах %.2f и %.2f, данные вашего прямоугольного треугольника будут следующими: периметр %.2f, площадь %.2f, гипотенуза %.2f", AB, AC, P, S, H)
}
