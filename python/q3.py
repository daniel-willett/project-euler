#!/usr/bin/python3

target=600851475143
#target=13195
def isPrime(n):
    for factor in range(2,int(n**0.5)+1):
        if n%factor==0:
            return False
    return True

for i in range(2,target+1):
    if target%i==0:
        if isPrime(target/i)==True:
            print(target/i)
            break
