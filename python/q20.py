#!/usr/bin/python3
import math

value = math.factorial(100)
arr = list(str(value))

total=0
for n in arr:
    total += int(n)

print(total)

