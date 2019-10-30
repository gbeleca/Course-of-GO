package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := big.NewInt(-1)
	b := big.NewInt(0)

	var limit big.Int
	limit.Exp(big.NewInt(10), big.NewInt(99), nil)

	for a.Cmp(&limit) < 0 {
		a.Add(a, b)
		a, b = b, a
	}
	fmt.Println(a)

}
