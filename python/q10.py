#!/usr/bin/python3

def isPrime(n):
    for i in range(2,int(n**0.5)+1):
        if n%i==0:
            return False
    return True
total=0
for i in range(2,2000000):
    if isPrime(i):
        total += i
print(total)
