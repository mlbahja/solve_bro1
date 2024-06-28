package main

import "fmt"

func main() {
	PrintMemory([10]byte{'h', 'e', 'l', 'l', 'o', 16, 21, '*'})
}

func PrintMemory(arr [10]byte) {
	arr1 := arr
	hex := "0123456789abcdef"
	strRev := ""
	count := 0
	idx := 0
	str := ""
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == 0 {
			str += "00 "
		} else {
			for arr1[i] > 0 {
				idx = int(arr1[i]) % 16
				strRev += string(hex[idx])
				arr1[i] /= 16
			}
			for i := len(strRev) - 1; i >= 0; i-- {
				str += string(strRev[i])
			}
			str += " "
			count++
			if count == 4 {
				str = str[:len(str)-1]
				str += "\n"
				count = 0
			}
			strRev = ""
		}
	}
	fmt.Println(str[:len(str)-1])
	for i := 0; i < len(arr); i++ {
		if arr[i] >= 32 && arr[i] <= 126 {
			fmt.Print(string(arr[i]))
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}
