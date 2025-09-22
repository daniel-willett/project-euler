#!/usr/bin/python3
n=100
sumSquared = (n*(n+1)/2)**2
squareSum=0
for i in range(1,n+1):
    squareSum+=i**2
print(sumSquared-squareSum)
