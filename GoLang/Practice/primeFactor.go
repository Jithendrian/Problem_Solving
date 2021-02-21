package main

import "fmt"

/*

The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?

 */

var(
	primeFactors[] int
	primeNumber int
)

func isPrime(a int) bool{

	for i:=2; i < a/2; i +=1 {
		if a%i == 0 {
			return false
		}
	}
	return true
}
func main() {
	primeNumber = 600851475143
	i := 2
	for primeNumber > 2 {
		if primeNumber % i == 0 && isPrime(i){
			primeNumber = primeNumber / i
			primeFactors = append(primeFactors, i)
		}
		i = i + 1
	}
	fmt.Println(primeFactors)
	fmt.Println(primeFactors[len(primeFactors) - 1])

}
