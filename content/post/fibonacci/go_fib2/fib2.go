package main

import (
	"fmt"
	"math/big"
)

func main() {
	const n = 2000000

	cur := big.NewInt(1)
	previous := big.NewInt(0)
	scratch := big.NewInt(0)
	fmt.Println(0)
	fmt.Println(cur)

	for i := 1; i < n; i++ {
		scratch.Add(cur, previous)
		cur, previous, scratch = scratch, cur, previous
	}

	fmt.Println(cur)
}
