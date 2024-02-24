#!/usr/bin/env python3
import sys
sys.set_int_max_str_digits(30000)
print(0)
f = 0
last = 1
print(last)
for i in range(1,100000):
    g = last
    last += f
    f = g
print(last)