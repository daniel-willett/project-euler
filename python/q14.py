#!/usr/bin/python3


def collatz(n):
    counter=0
    while n>1:
        if n%2==0:
            n=n/2
        else:
            n=(3*n)+1
        counter+=1
    return counter

runningN=0
runningCollatz=0
n=1
while n<1000000:
    value=collatz(n)
    if value>runningCollatz:
        runningCollatz=value
        runningN=n
    n+=1
print(runningN)
