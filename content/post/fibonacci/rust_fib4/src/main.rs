use num_bigint::BigUint;
use num_traits::{One, Zero};

// Translated from https://www.nayuki.io/res/fast-fibonacci-algorithms/FastFibonacci.cs
// Increased iterations to many millions

fn main() {
    const N: u32 = 100000000;
    let a = fib(N);
    println!("{}", a.bits());
}

fn fib(n: u32) -> BigUint {
    let mut a = BigUint::zero();
    let mut b = BigUint::one();
    for i in (0..32).rev() {
        let d = &a * (&b * 2u32 - &a);
        let e = &a * &a + &b * &b;
        a = d;
        b = e;
        if (n >> i) & 1 != 0 {
            a += &b;
            std::mem::swap(&mut a, &mut b);
        }
    }
    a
}
