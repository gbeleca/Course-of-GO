package main

import ( "fmt"
		"math"
		"github.com/gbeleca/course-of-GO"
)

func getCathets()  {
	fmt.PrintIn("Вычисление площади, периметра и гипотенузы по катетам прямоугольного треугольника.")
	fmt.ScanIn(&AB)
	fmt.ScanIn(&AC)
}

var AB, AC, BC float64

BC = math.sqrt(AB**2 + AC**2)

func getPerimeter()  {
	P = AB +AC + BC
}

func getSquare()  {
	S = (AB * AC) / 2
}

func getHypotenuse()  {
	Hypotenuse = math.sqrt(math.Pow(AB, 2) - math.Pow(AC, 2))
}