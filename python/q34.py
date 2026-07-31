#!/usr/bin/python3

import math

def factSum(n):
    runningTotal = 0
    for digit in str(n):
        runningTotal += math.factorial(int(digit))
    return runningTotal


#I don't know an upper bound so just going to guess 100,000 and hope the program doesn't take too long
total = 0
for i in range(3,100000):
    if factSum(i)==i:
        total += i
print(total)
