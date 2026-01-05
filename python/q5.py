#!/usr/bin/python3

"""
We take each number 1-20 and factor it into it's prime decomposition
Then we take each power for each value of the decomposition and check if it's larger than a power we already have in a counter for those primes
In the end we should end up with a number divisible by each 1-20 by construction, but doing this minimally so by being divisible just enough
"""

counter = {2:0,
           3:0,
           5:0,
           7:0,
           11:0,
           13:0,
           17:0,
           19:0}

#All the primes and the number of times we've seen the counted for (initially this is zero)

def factor(n):
    primes = [2,3,5,7,11,13,17,19] #We are only factoring numbers upto 20
    result=[]
    for p in primes:
        cannotDivide = False
        numOfDivisors = 0
        while cannotDivide == False:
            if n%p!=0:
                cannotDivide = True
            else:
                numOfDivisors += 1
                n = n/p
        result.append([p,numOfDivisors])
    return result


for i in range(2,20+1):
    factors = factor(i)

    for valuePair in factors:
        prime = valuePair[0]
        if counter[prime]<valuePair[1]:
            counter[prime] = valuePair[1]
#Now we have the primes and their min-max power so it's just time to combine into one number
total = 1
for prime in counter:
    total = total * (prime ** counter[prime])

print(total)
