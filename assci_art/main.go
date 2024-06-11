package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Message Error")
		os.Exit(1)
	}
	rawData, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Message Error")
		return
	}
	asciiTable := asciiTableMaker(string(rawData))
	arg := strings.ReplaceAll(os.Args[1], "\\n", "\n")
	arr := strings.Split(arg, "\n")
	for _, str := range arr {
		for _, char := range str {
			if char < ' ' || char > '~' {
				fmt.Println("message error  character out of range")
				os.Exit(1)
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
func asciiTableMaker(rawData string) [][]string {
	dataTmp := strings.Split(rawData, "\n")
	tmp := []string{}
	asciiTable := [][]string{}
	for _, l := range dataTmp {
		if len(l) > 0 {
			tmp = append(tmp, l)
		} else {
			if len(tmp) > 0 {
				asciiTable = append(asciiTable, tmp)
				tmp = []string{}
			}
		}
	}
	return asciiTable
}
