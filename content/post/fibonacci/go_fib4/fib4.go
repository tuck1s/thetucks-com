package main

// Translated from https://www.nayuki.io/res/fast-fibonacci-algorithms/FastFibonacci.cs
// Increased iterations to many millions

import (
	"fmt"
	"math/big"
)

func main() {
	const n = 100000000
	a := fib(n)
	fmt.Println(a.BitLen())
}

func fib(n uint) *big.Int {
	a := big.NewInt(0)
	b := big.NewInt(1)

	// Preallocate temp vars
	temp := new(big.Int)
	temp2 := new(big.Int)
	temp3 := new(big.Int)

	for i := 31; i >= 0; i-- {
		// Calculate d = a * (2*b - a) .. Tried (b+b) as well, seems about the same speed
		temp.Mul(b, big.NewInt(2))
		d := temp.Sub(temp, a)
		d.Mul(d, a)

		// Calculate e = a^2 + b^2 and assign to b
		temp2.Mul(a, a)
		temp3.Mul(b, b)
		b.Add(temp2, temp3)

		a.Set(d)
		if (uint(n)>>uint(i))&1 != 0 {
			a.Add(a, b)
			a, b = b, a
		}
	}
	return a
}
