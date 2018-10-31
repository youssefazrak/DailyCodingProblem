
/*
Given an array of integers, find the first missing positive integer in linear
time and constant space. In other words, find the lowest positive integer that
does not exist in the array. The array can contain duplicates and negative numbers as well.

For example, the input [3, 4, -1, 1] should give 2. The input [1, 2, 0] should give 3.

You can modify the input array in-place.
*/

package main

import (
	"fmt"
	"sort"
)

func main() {
	t := []int{3, 4, -1, 1, 0, 0, 1, -2, -4, -2, 7}
	l := []int{1, 2, 3, 4, 0}

	sort.Ints(t)
	sort.Ints(l)

	var lowest_num int = 1

	for _, num := range t {
		if num > 0 {
			if num > lowest_num {
				break
			} else {
				lowest_num++
			}
		}
	}
	fmt.Print(lowest_num)
}


