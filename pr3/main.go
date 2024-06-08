/*
func romanToInt(s string) {

	Iterate over the string MCMXCIV:

Start with the first character M (1000):

Next character is C (100), which is less than M.
Add 1000 to total: total = 1000.
Next character C (100):

Next character is M (1000), which is greater than C.
Subtract 100 from total: total = 1000 - 100 = 900.
Next character M (1000):

Next character is X (10), which is less than M.
Add 1000 to total: total = 900 + 1000 = 1900.
Next character X (10):

Next character is C (100), which is greater than X.
Subtract 10 from total: total = 1900 - 10 = 1890.
Next character C (100):

Next character is I (1), which is less than C.
Add 100 to total: total = 1890 + 100 = 1990.
Next character I (1):

Next character is V (5), which is greater than I.
Subtract 1 from total: total = 1990 - 1 = 1989.
Last character V (5):

No next character.
Add 5 to total: total = 1989 + 5 = 1994.

	romain := []rune(s)

	fmt.Println(romain)
	fmt.Println(str)

}

	func main() {
		romanToInt("III")
		fmt.Println(romanToInt("III"))
		fmt.Println(romanToInt("III"))

}

'I', 'V', 'X', 'L', 'C', 'D', 'M')
----------------------------------
'I', 'V', 'X', 'L', 'C', 'D', 'M')
I >>>>>>>>>>> 1
V >>>>>>>>>>> 5
X >>>>>>>>>>> 10
L >>>>>>>>>>> 50
C >>>>>>>>>>> 100
D >>>>>>>>>>> 500
M >>>>>>>>>>> 1000
*/
package main

import (
	"fmt"
)

func romanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	fmt.Println(len(romanMap))
	total := 0
	//n := len(s)
	for i := 0; i < 7; i++ {
		currentValue := romanMap[s[i]]
		fmt.Println(currentValue)
		if i+1 < 7 && romanMap[s[i+1]] > currentValue {
			total -= currentValue
		} else {
			total += currentValue
		}
	}
	return total
}

func main() {
	fmt.Println(romanToInt("MCMXCIV"))
	// Output: 1994
}
