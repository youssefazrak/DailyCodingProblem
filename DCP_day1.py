#!/usr/bin/python3
from random import randint

#k = randint(0,28)
#print(k)

b = randint(1,25) 
print('k is ', b)
a = [4,3,10,15,5,7,1,2,4,9,8,7]
print(a)
print('length of the list is', len(a))
#if a.index(0) == False:
#    print('There is no zero in the list')

#if there is a number greater than b then we remove it from the list
#taking into account that we don't have negative numbers
for i in range(0, len(a) - 1):
    for j in range(i + 1, len(a)):
        #print('first value is ', a[i], ' second value is ', a[j])
        if a[i] + a[j] == b:
            print('Nicely done!')
            print('first value is ', a[i], ' second value is ', a[j])
        #else:
         #   print('There are no two numbers that sums up to ', b)
            
