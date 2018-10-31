#!/bin/bash/python3

"""
Given an array of integers, find the first missing positive integer in linear
time and constant space. In other words, find the lowest positive integer that
does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
"""

t = [3, 4, -1, 1, 0, 0, 1, -2, -4, -2, 7]
l = [1, 2, 3, 4, 0] 


def lowest(lis):
    lis.sort()
    lowest_num = 1

    for i in lis:
        if i > 0:
            if i > lowest_num:
                break
            else:
                lowest_num = i + 1
    print(lowest_num)

lowest(t)
lowest(l)
       



