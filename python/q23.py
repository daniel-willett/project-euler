#!/usr/bin/python3

def divisors(n):
    lst = [1]
    for i in range(2,int(n**0.5)):
        if n%i==0:
            lst.append(i)
            lst.append(n/i)
    if n**0.5==int(n**0.5):
        lst.append(n**0.5)
    return lst

def isAbundant(n):
    divs = divisors(n)
    total = sum(divs)
    if total>n:
        return True
    return False

abundantNumbers = []
for i in range(12,28123+1):
    if isAbundant(i):
        abundantNumbers.append(i)

print(len(abundantNumbers))

abundantSum = []
for num1 in abundantNumbers:
    for num2 in abundantNumbers:
        abundantSum.append(num1+num2)

total = 0
for i in range(12,28123+1):
    if not i in abundantSum:
        total += i

print(total)
