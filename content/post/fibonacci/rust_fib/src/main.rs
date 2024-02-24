use num_bigint::BigUint;
use num_traits::Zero;

fn main() {
    const N: usize = 2000000;

    let mut last = BigUint::from(1u32);
    let mut f = BigUint::zero();
    let zero = BigUint::zero();

    println!("{}", zero);
    println!("{}", last);

    for _ in 1..N {
        let g = last.clone();
        last += &f;
        f = g;
    }

    println!("{}", last);
}
