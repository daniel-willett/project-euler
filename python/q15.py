#!/usr/bin/python3

"""
We have 20x20 grid. 
We can only do R(ight) or D(own)
They are not dependant on each other and so can be interchanged.
So every path is a sequence of 40 characters: 20 R's and 20 D's
And they can be changed around freely.
i.e. solution = 40!/(20!*20!)
"""
import math as m
print(m.factorial(40)/(m.factorial(20)*m.factorial(20)))
