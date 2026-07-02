#!/usr/bin/python3

units = {0:"",
         1:"one",
         2:"two",
         3:"three",
         4:"four",
         5:"five",
         6:"six",
         7:"seven",
         8:"eight",
         9:"nine",
         10:"ten",
         11:"eleven",
         12:"twelve",
         13:"thirteen",
         14:"fourteen",
         15:"fifteen",
         16:"sixteen",
         17:"seventeen",
         18:"eighteen",
         19:"nineteen"}
#10 to 19 don't play well in english so have to catch them separately
tens = {2:"twenty",
        3:"thirty",
        4:"forty",
        5:"fifty",
        6:"sixty",
        7:"seventy",
        8:"eighty",
        9:"ninety"}
hundred = "hundred"

def numberToWord(n):
    theRest=n
    word=""
    if theRest//100!=0:
        word += units[theRest//100] + hundred
        theRest = theRest%100
        if theRest!=0:
            word += "and"
    if theRest>19:
        word += tens[theRest//10] + units[theRest%10]
    elif theRest!=0:
        word += units[theRest] 
    return word


string = ""
for num in range(1,1000):
    string += numberToWord(num)

print(string)
print(len(string)+len("OneThousand"))

"""
ten
eleven
twelve
thirteen
fourteen
fifteen
sixteen
seventeen
eighteen
nineteen
twenty
twenty one
twenty two
...
twenty nine
thirty
thirty one
thirty two
...
ninety eight
ninety nine
one hundred
one hundred and one
one hundred and two
...
one hundred and ninety nine
two hundred
two hundred and one
...
nine hundred and ninety nine
one thousand
"""
