#!/usr/bin/python3
for c in range(1000,1,-1):
    for b in range(1000-c,1,-1):
        for a in range(b,1,-1):
            if (a+b+c==1000):
                if (a**2 + b**2 == c**2):
                    print(a*b*c)
