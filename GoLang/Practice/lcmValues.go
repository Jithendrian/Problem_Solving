package main

/*
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?

 */

import "fmt"

func lcm(a int, b int) int {
	var greater int = 0
	if a > b{
		greater = a
	}else{
		greater = b
	}
	for true{
		if (greater % a == 0 && greater %b == 0){
			return greater
		}
		greater +=1
	}
	return greater
}

func main() {
	lcmVal := 1
	for i:=2; i<21; i +=1{
		lcmVal = lcm(lcmVal, i)
	}
	fmt.Println(lcmVal)
}
