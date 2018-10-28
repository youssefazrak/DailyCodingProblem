#!/usr/bin/python3

"""
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
Bonus: Can you do this in one pass?

"""

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
            
