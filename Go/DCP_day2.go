/*
Given an array of integers, return a new array such that each element at index i of the new array
is the product of all the numbers in the original array except the one at i.

For example, if our input was [1, 2, 3, 4, 5], the expected output would be [120, 60, 40, 30, 24]. 
If our input was [3, 2, 1], the expected output would be [2, 3, 6].

Follow-up: what if you can't use division?

This solution provides no division. At this point I haven't found a proper solution to use with division.
Probably related to the length that you divide.
In all cases this one seems clean. We basically separate the list into two: one before the index and 
one after the index. We multiple the elements of the first list between themselves till the index minus one. 
We do the same for the second list. In the end we multiply both results together and we have the final results.
*/


package main

import "fmt"

func main() {
	a := [5]int{ 1, 2, 3, 4, 5 }
	b := [5]int{}

	for i:= 0; i < len(a); i++ {
		v := 0
		v1 := 1
		v2 := 1

		for j := 0; j < i; j++ {
			v1 *= a[j]
		}

		for j := i + 1; j < len(a); j++ {
			v2 *= a[j]
		}

		v = v1 * v2
		b[i] = v // another option to try is "b := append(b, v)"

	}
	fmt.Println(b)
}






















