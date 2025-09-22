#!/usr/bin/python3
def isPalendrome(n):
    return str(n)[::-1]==str(n)

def program():
    counter=[]
    for i in range(999,100,-1):
        for j in range(999,100,-1):
            n=i*j
            if isPalendrome(n):
                counter.append(n)
    return max(counter)
print(program())
