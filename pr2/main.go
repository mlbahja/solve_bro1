package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	n := x
	k := 0
	for n > 0 {
		k *= 10
		k += n % 10
		n /= 10

	}
	if x == k {
		return true
	}
	return false

}
func revers(s string) bool {
	result := ""
	for i := len(s) - 1; i >= 0; i-- {
		result += string(s[i])
	}
	// if result == s {
	// 	return true
	// }
	// return false
	return result == s

}
func main() {
	fmt.Println(isPalindrome(1221))
	fmt.Println(revers("tinit"))
}
