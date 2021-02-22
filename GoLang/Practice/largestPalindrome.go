package main

import (
	"fmt"
	"sort"
)

/*
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers
 */

var(
	combinations = make(map[int]int)
)
func powerCalculation(a int, b int) int {
	pow := 1
	for i:=1 ; i <= b; i +=1{
		pow = pow * a
	}
	return pow
}

func isPalindrome(num int) bool{
	factor := 10
	var digits[] int
	var trailing bool = true
	var origNum int = num
	for num > 0 {
		if(num % factor) == 0 && trailing {
			num = num / factor
		}else{
			trailing = false
			digits = append(digits, (num % factor))
			num = num / factor
		}
	}
	i := 0
	revNum := 0
	for i < len(digits){
		revNum = revNum + digits[i] * powerCalculation(factor, (len(digits) - i - 1))
		i += 1
	}
	if origNum == revNum{
		return true
	}
	return false
}

func main() {
	var palindromes[] int
	m := 999
	for m <= 999 && m > 99{
		j := 999
		for j <= 999 && j > 99{
			if isPalindrome(m*j) {
				palindromes = append(palindromes, m*j)
			}
			j = j - 1
		}
		m = m - 1
	}
	sort.Ints(palindromes)
	fmt.Println("Largest palindrome number %d", palindromes[len(palindromes) - 1])
}
