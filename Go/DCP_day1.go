
/*
Given a list of numbers and a number k, return whether any two numbers from the list add up to k.
For example, given [10, 15, 3, 7] and k of 17, return true since 10 + 7 is 17.
Bonus: Can you do this in one pass?
*/

package main

import "fmt"


func main() {
	k := 15 //random number
	a := [10]int{ 4, 5, 2, 1, 7, 8, 10, 6, 12, 3 } //random list
	fmt.Println("k value is: ", k)
	fmt.Println("the list is: ", a)

	for i := 0; i < len(a) - 1; i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] + a[j] == k {
				fmt.Println("Nicely done!")
				fmt.Println("First value is: ", a[i], " Second value is: ", a[j])
			}
		}
	}
}

