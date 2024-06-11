package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	if len(args) != 2 {
		fmt.Println("message error")
		os.Exit(1)
	}
	if args[1] != "standard" && args[1] != "shadow" && args[1] != "thinkertory" {
		fmt.Println("Usage: ascii_art [string]")
		os.Exit(1)
	}
	rawData, err := os.ReadFile(args[1] + ".txt")
	if err != nil {
		fmt.Println("message error")
		return
	}
	//thinkort file contient \n\r
	replace := strings.ReplaceAll(string(rawData), "\r", "")
	asciiTable := asciiTableMaker(string(replace))
	arg := strings.ReplaceAll(args[0], "\\n", "")
	arr := strings.Split(arg, "\n")

	for _, str := range arr {
		for _, char := range str {
			if char < ' ' || char > '~' {
				fmt.Println("message error character out of range ")
				return
			}
		}
	}
	if len(arg) == strings.Count(arg, "\n") {
		arr = arr[:len(arr)-1]
	}
	for _, str := range arr {
		if len(str) == 0 {
			fmt.Println()
		} else {
			for i := 0; i < 8; i++ {
				for _, char := range str {
					if char >= ' ' && char <= '~' {
						fmt.Print(asciiTable[int(char)-int(' ')][i])
					}
				}
				fmt.Println()
			}
		}
	}
}
func asciiTableMaker(rowData string)[][]string{
	datatmp := strings.Split(rowData, "\n")
	tmp := []string{}
	asciiTable := [][]string{}

	for _, char := range datatmp{
		if len(char) > 0 {
			tmp = append(tmp, char)
		}else {
			if len(tmp) > 0 {
				asciiTable = append(asciiTable, tmp)
				tmp = []string{}
			}
		}
	}
	return asciiTable 
}
