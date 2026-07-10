#!/usr/bin/python3

"""
Start bottom up.
If we have a set of two rows e.g.

    3  1  12 9 10
  10 34 28 12 59 39
 3  2  5  9  9  2  1

Then we only look at the possible paths and reduce by half by taking the larger path. Take the first section for example,

  10
 3  2

We know the max path can't be 10+2 because there is definitely a larger path by picking 10+3. Then we move on and apply the same to the next part:
  34
 2  5
Same again, we see that 5>2 so we'd pick 34+5. After repeating this all the way along we end up with

  10 34 28 12 59 39
  becomes
  13 39 37 21 68 41
Now we do this all over again using this new calculated row to be the 'bottom' i.e.
    3  1  12 9 10
  13 39 37 21 68 41
"""

#Note to self, can get rid of leading 0's in vim by doing `:%s/\(\W\)0\(\d\)/\1\2/g`

arr=[[75],
    [95,64],
    [17,47,82],
    [18,35,87,10],
    [20,4,82,47,65],
    [19,1,23,75,3,34],
    [88,2,77,73,7,63,67],
    [99,65,4,28,6,16,70,92],
    [41,41,26,56,83,40,80,70,33],
    [41,48,72,33,47,32,37,16,94,29],
    [53,71,44,65,25,43,91,52,97,51,14],
    [70,11,33,28,77,73,17,78,39,68,17,57],
    [91,71,52,38,17,14,91,43,58,50,27,29,48],
    [63,66,4,68,89,53,67,30,73,16,69,87,40,31],
    [4,62,98,27,23,9,70,98,73,93,38,53,60,4,23]]


for row in range(len(arr)-1,-1,-1):
    for pos in range(len(arr[row])-1):
        if arr[row][pos]>arr[row][pos+1]:
            arr[row-1][pos] += arr[row][pos]
        else:
            arr[row-1][pos] += arr[row][pos+1]

print(arr[0:1])
