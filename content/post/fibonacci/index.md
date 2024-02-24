---
title: 'Generating the Fibonacci sequence on a Raspberry Pi 5'
date: 2024-02-24T11:50:03Z
tags: ['arithmetic', 'linux', 'raspberry pi']

---
Linux includes a bench calculator called [`bc`](https://linux.die.net/man/1/bc), which can achieve interesting arbitrary-precision results in very few lines of "code". For ease, this can be embedded in a [here document](https://linux.die.net/abs-guide/here-docs.html):

_fib.sh_
```bash
#!/usr/bin/env bash
BC_LINE_LENGTH=0 bc -q <<end
0;1; for (i=1;i<1000;i++) {g=last;last+=f;f=g;last}
end
```

The `BC_LINE_LENGTH` environment variable is used to prevent long lines from being split with a `\` continuation character.

The `-q` flag suppresses the `>>>` welcome prompt.

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

Here's the same routine in Python:

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

This produces identical output, because Python integers are also arbitrary precision. Python is significantly slower (in this case, Python 3.11.2 on a Raspberry Pi 5):

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

On longer runs such as 100 000, Python gives an error:

```
ValueError: Exceeds the limit (4300 digits) for integer string conversion; use sys.set_int_max_str_digits() to increase the limit
```

Changing this setting to 30000 allows the program to run.

Both languages must spend less time calculating the numbers, and more time serializing the huge integers to strings - not suprrising whn the final result is 20899 digits long! In particular Python `print( )` output is costly.

Changing both programs to only print the first two and the final number, shows an interesting result: Python is _faster_ over longer runs.

```
time ./fib.sh >/dev/null ; time ./fib.py >/dev/null

real    0m2.488s
user    0m2.478s
sys     0m0.008s

real    0m0.589s
user    0m0.580s
sys     0m0.008s
```

