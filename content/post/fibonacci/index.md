---
title: 'Generating the Fibonacci sequence on a Raspberry Pi'
date: 2024-02-24T11:50:03Z
tags: ['arithmetic', 'linux', 'raspberry pi']

---
Linux includes an arbitrary-precision calculator called [`bc`](https://linux.die.net/man/1/bc) which can achieve interesting results in very few lines. For ease, this can be embedded in a [here document](https://linux.die.net/abs-guide/here-docs.html):

_fib.sh_
```bash
#!/usr/bin/env bash
BC_LINE_LENGTH=0 bc -q <<end
0;1; for (i=1;i<1000;i++) {g=last;last+=f;f=g;last}
end
```

Arbirtary precision is very useful for generating long numbers accurately.

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