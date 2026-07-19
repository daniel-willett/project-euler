#!/usr/bin/python3

F1=1
F2=1
temp=0
index=2
while F2//(10**999)==0:
    temp = F1+F2
    F1=F2
    F2=temp
    index += 1
print(F1)
print(F2)
print(index)
