package main

import (
	"fmt"
	"math/big"
)

func main() {
	const n = 2000000

	last := big.NewInt(1)
	f := big.NewInt(0)
	fmt.Println(big.NewInt(0))
	fmt.Println(last)

	for i := 1; i < n; i++ {
		g := new(big.Int).Set(last)
		last.Add(last, f)
		f.Set(g)
	}

	fmt.Println(last)
}
