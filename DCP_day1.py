#!/usr/bin/python3
from random import randint

k = randint(1,25) 
print('k is ', k)

a = [4,3,10,15,5,7,1,2,4,9,8,7]
print(a)
print('length of the list is', len(a))

for i in range(0, len(a) - 1):
    for j in range(i + 1, len(a)):
        if a[i] + a[j] == k:
            print('Nicely done!')
            print('first value is ', a[i], ' second value is ', a[j])
            
