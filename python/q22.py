#!/usr/bin/python3

def productOfChars(string):
    counter = 0
    for char in string:
        counter += ord(char)-64
    return counter

f = open("q22_names.txt")
val = f.read()
val = val.replace("\"", "")
nameList = val.split(",")
nameList = sorted(nameList)


val=0
for i in range(len(nameList)):
    val += productOfChars(nameList[i])*(i+1)

print(val)
