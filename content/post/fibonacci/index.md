---
title: 'Generating the Fibonacci sequence on a Raspberry Pi 5'
date: 2024-02-24T11:50:03Z
tags: ['arithmetic', 'linux', 'raspberry pi']

---
Linux includes a bench calculator called [`bc`](https://linux.die.net/man/1/bc), which can achieve remarkable arbitrary-precision results in very few lines of code. For ease of running, code can be embedded in a [here document](https://linux.die.net/abs-guide/here-docs.html):

_fib.sh_
```bash
#!/usr/bin/env bash
BC_LINE_LENGTH=0 bc -q <<end
0;1; for (i=1;i<1000;i++) {g=last;last+=f;f=g;last}
end
```

The `BC_LINE_LENGTH` environment variable is used to prevent long lines from being split with a `\` continuation character.

The `-q` flag suppresses the welcome prompt.

The `0;1;` bit just prints the first two results. The special variable `last` will have the value `1` after this.

The `for` loop is C-like in its construction, iterating `i` over the specified range.

Inside the `{` `}` braces, variable `last` is updated, with variables `f` and `g` carrying the previous results forward.

The output's first few lines look like this; we can time how long it takes.

```bash
time ./fib.sh
0
1
1
2
3
5
8
13
21
34
55
:
:

real    0m0.034s
user    0m0.001s
sys     0m0.019s
```

`bc` has many of the features of larger programming languages such as loops and functions.

Here's the same routine in Python, producing identical output, because Python integers are also arbitrary precision:

```python
#!/usr/bin/env python3
print(0)
f = 0
last = 1
print(last)
for i in range(1,1000):
    g = last
    last += f
    f = g
    print(last)
```

Python is significantly slower (Python 3.11.2 on a Raspberry Pi 5):

```
real    0m0.069s
user    0m0.051s
sys     0m0.016s
```

On longer runs (such as 10 000 iterations) the difference is even clearer. Here we direct console output to /dev/null to reduce the effect of printing/scrolling the output.

```
 time ./fib.sh >/dev/null ; time ./fib.py >/dev/null

real    0m0.252s
user    0m0.243s
sys     0m0.008s

real    0m12.409s
user    0m12.391s
sys     0m0.004s
```

On longer runs such as 100000, Python gives an error:

```
ValueError: Exceeds the limit (4300 digits) for integer string conversion; use sys.set_int_max_str_digits() to increase the limit
```

Changing this setting to 30000 allows the program to run.

### What makes Python slower or faster than bc?

As the numbers become longer, both languages spend more time serializing them to strings. This is unsurprising when fib(100000) is 20899 digits long! In particular Python `print( )` output appears to be costly. Python is also a much larger binary program to load and start (on my Pi 5, python3 = 5MB vs bc = 75KB).

Changing both programs to only print 0, 1 and the final number shows an interesting result: Python becomes _significantly faster_ over longer iterations:

```
time ./fib.sh >/dev/null ; time ./fib.py >/dev/null

real    0m2.488s
user    0m2.478s
sys     0m0.008s

real    0m0.589s
user    0m0.580s
sys     0m0.008s
```

More details on the Pi 5 system:
```
uname -a
Linux pi5 6.1.0-rpi4-rpi-v8 #1 SMP PREEMPT Debian 1:6.1.54-1+rpt2 (2023-10-05) aarch64 GNU/Linux
```

### On a more powerful computer

On a Macbook Pro M3, 100000 iterations are very fast (5x and 7x faster than the Pi 5).

```
time ./fib.sh >/dev/null ; time ./fib.py >/dev/null
./fib.sh > /dev/null  0.32s user 0.01s system 99% cpu 0.334 total
./fib.py > /dev/null  0.11s user 0.01s system 79% cpu 0.146 total
```

```
uname -a
Darwin SteveT-M3.local 23.2.0 Darwin Kernel Version 23.2.0: Wed Nov 15 21:54:51 PST 2023; root:xnu-10002.61.3~2/RELEASE_ARM64_T6030 arm64
```

Here are the final [fib.py](./fib.py) and [fib.sh](./fib.sh) programs.

### What next?

Go and Rust languages have reputations for being fast (and faster). Both have arbitrary-precision arithmetic packages.

Go
```
time ./fib >/dev/null
./fib > /dev/null  0.07s user 0.01s system 115% cpu 0.071 total
```

Rust:
```
 time ./target/release/rust_fib >/dev/null
./target/release/rust_fib > /dev/null  0.08s user 0.00s system 97% cpu 0.083 total
```

Increasing the iterations to 2 million:

Go:
```
go build
time ./go_fib >/dev/null
./go_fib > /dev/null  19.50s user 2.93s system 122% cpu 18.363 total
```

Rust:
```
cargo build --release
time ./target/release/rust_fib >/dev/null
./target/release/rust_fib > /dev/null  27.81s user 0.02s system 99% cpu 27.866 total
```

Both [go](go_fib/fib.go) and [rust](rust_fib/src/main.rs) programs produce [identical](go_fib/go_fib_2M.txt) [output](rust_fib/rust_fib_2M.txt). Fib(2000000) is over 400000 digits long.

### The Rustacean claws it back

I just realised I had x64 binaries left over on my system, as I migrated it from an older machine.

Now got the arm64 binaries, and checking the built binary is also arm64:

```
cargo build --release
    Finished release [optimized] target(s) in 0.00s

% time ./target/release/rust_fib >rust_fib_2M_new.txt
./target/release/rust_fib > rust_fib_2M_new.txt  19.74s user 0.01s system 98% cpu 19.966 total

% file ./target/release/rust_fib
./target/release/rust_fib: Mach-O 64-bit executable arm64
```

With an improvement to the inner loop from [r/rust](https://www.reddit.com/r/rust/comments/1az4tfn/comment/krzgtw1/?utm_source=share&utm_medium=web2x&context=3), Rust is the front-runner.

```
time ./target/release/rust_fib >/dev/null
./target/release/rust_fib > /dev/null  16.85s user 0.01s system 99% cpu 16.885 total
```

Toolchain:
```
cargo 1.76.0 (c84b36747 2024-01-18)
rustc 1.76.0 (07dca489a 2024-02-04)
```

### Going further

A similar inner-loop improvement from [r/golang](https://www.reddit.com/r/golang/comments/1az4r01/fibonacci_sequences_using_bc_python_go_and_rust/) reclaims the top spot:

```
go build -ldflags "-s -w"

time ./go_fib2 >go_fib2_2M.txt
./go_fib2 > go_fib2_2M.txt  5.79s user 0.03s system 97% cpu 5.996 total
```

This version is [go_fib2](go_fib2/fib2.go).

Toolchain:
```
go version go1.22.0 darwin/arm64
```

### Improve the algorithm!

[This post](https://www.reddit.com/r/golang/comments/1az4r01/comment/ks0pj2s/?utm_source=share&utm_medium=web2x&context=3) uses the much better "fast doubling" algorithm in Go (and importantly, gives the same output).

This version is [go_fib3](go_fib3/fib3.go) and is amazingly fast on large numbers - the paper below states Θ(log n) complexity.

```
 time ./go_fib3 >/dev/null
./go_fib3 > /dev/null  0.09s user 0.01s system 57% cpu 0.178 total
```

Here's a naive translation of this algorithm into Rust: [rust_fib3](rust_fib3/src/main.rs), currently slower than Go.

```
time ./target/release/rust_fib3 >/dev/null
./target/release/rust_fib3 > /dev/null  0.37s user 0.00s system 73% cpu 0.506 total
```

```
 time ./go_fib3 >/dev/null
./go_fib3 > /dev/null  0.09s user 0.01s system 57% cpu 0.178 total
```

### Finally

Starting from the C# version on [Nayuki's page](https://www.nayuki.io/page/fast-fibonacci-algorithms) made further gains possible.

* [go_fib4](go_fib4/fib4.go)

* [rust_fib4](rust_fib4/src/main.rs)

Printing the number to the console in ASCII was dominating the run time, even when output was directed to `/dev/null`, so these versions print the number's length (in bits) instead.

To show up speed differences, these final versions calculate the 100 millionth number.

```
time ./go_fib4
69424191
./go_fib4  11.10s user 0.03s system 99% cpu 11.205 total
```

```
time ./target/release/rust_fib4
69424191
./target/release/rust_fib4  11.54s user 0.04s system 98% cpu 11.737 total
```

### Further reading

1. [Nayuki](https://www.nayuki.io/page/fast-fibonacci-algorithms)'s page on fast Fibonacci algorithms, referred to via the Reddit reply linked above.