package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filePath string = "./input.txt"

func parseInput(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func parseCommand(s string) (cmd string, x1, y1, x2, y2 int) {

	return "toggle", 0, 0, 0, 0
}

func main() {
	//parseInput(filePath)

	test1 := "toggle 628,958 through 811,992"
	fmt.Println(strings.Index(test1, "toggle"))
}
