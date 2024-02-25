use num_bigint::BigUint;
use num_traits::Zero;
use num_traits::One;

fn main() {
    const N: u32 = 2000000;
    let (a, _) = fib(N);
    println!("{}", a);
}

fn fib(n: u32) -> (BigUint, BigUint) {
    if n == 0 {
        (BigUint::zero(), BigUint::one())
    } else {
        let (a, b) = fib(n / 2);
        let c = &a * (&b * 2u32 - &a);
        let d = &a * &a + &b * &b;
        if n % 2 == 0 {
            (c, d)
        } else {
            // Improvement from https://www.reddit.com/r/rust/comments/1az4tfn/comment/ks3j165/?utm_source=share&utm_medium=web2x&context=3
            let e = c + &d;
            (d, e)
        }
    }
}
