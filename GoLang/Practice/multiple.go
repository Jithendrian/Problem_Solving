package main

import "fmt"

//If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

//Find the sum of all the multiples of 3 or 5 below 1000.

func main() {
	var x int
	x = 1
	var multiples[] int
	for x < 1000 {
		if x % 3 == 0 || x % 5 == 0 {
			multiples = append(multiples, x)
		}
		x += 1
	}
	fmt.Println(multiples)
	var sum int
	for i:=0; i< len(multiples); i = i+1 {
		sum = sum + multiples[i]
	}
	fmt.Println(sum)
}
