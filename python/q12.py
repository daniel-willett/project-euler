#!/usr/bin/python3

def numOfDivs(x):
    counter=0
    for i in range(1,int(x**0.5)+1):
        if x%i==0:
            counter+=1
    counter*=2
    if x==x**0.5:
        counter-=1
    return counter
n=1
triangle=1
while numOfDivs(triangle)<=500:
    n+=1
    triangle=n*(n+1)/2
print(triangle)
