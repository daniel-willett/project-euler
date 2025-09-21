#!/usr/bin/python3

arr=[1,2]
counter=1
total=0
while arr[counter]<=4000000:
    if arr[counter]%2==0:
        total+=arr[counter]
    arr.append(arr[counter]+arr[counter-1])
    counter+=1
print(total)
