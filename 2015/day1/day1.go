package main

import (
	"bufio"
	"fmt"
	"os"
)

const filePath = "./input.txt"

func countParens(input string) int {
	ct := 0
	for _, c := range input {
		if c == '(' {
			ct++
		} else if c == ')' {
			ct--
		} else {
			//do nothing on invalid char
		}
	}
	return ct
}

func enterBasement(input string) int {
	ct := 0
	for i, c := range input {
		if c == '(' {
			ct++
		} else if c == ')' {
			ct--
		} else {
			//do nothing on invalid char
		}
		if ct < 0 {
			return i + 1
		}
	}
	return -1
}

func readFile(filePath string) string {
	inFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	reader := bufio.NewScanner(inFile)
	var outString string
	for reader.Scan() {
		outString = outString + reader.Text()
	}
	return outString
}

func main() {
	//fmt.Println(readFile(filePath))

	instring := readFile(filePath)
	//destFloor := countParens(instring)
	fmt.Println(countParens(instring))
	fmt.Println(enterBasement(instring))
}
