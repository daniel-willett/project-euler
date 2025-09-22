#!/usr/bin/python3

def isPrime(n):
    for i in range(2,int(n**0.5)+1):
        if n%i==0:
            return False
    return True
counter=2
num=4
while counter<10001:
    if isPrime(num):
        counter += 1
    num+=1
print(num-1)
