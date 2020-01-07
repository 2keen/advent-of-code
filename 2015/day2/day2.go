package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const filePath = "./input.txt"

func readFile(filePath string) (int, int) {
	inFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()
	reader := bufio.NewScanner(inFile)
	paper := 0
	ribbon := 0
	for reader.Scan() {
		sides := strings.Split(reader.Text(), "x")
		lwh := make([]int, len(sides))
		for i, n := range sides {
			num, err := strconv.Atoi(n)
			if err != nil {
				panic(err)
			}
			lwh[i] = num
		}

		paper += paperNeeded(lwh)
		ribbon = ribbon + ribbonNeeded(sort(lwh))
		fmt.Println(lwh, sort(lwh), ribbonNeeded(sort(lwh)), "\t", ribbon)
	}
	return paper, ribbon
}

func sort(array []int) []int {
	arrayLength := len(array)
	outArray := make([]int, arrayLength)

	for _, n := range array {
		ct := 0
		place := 0
		for _, c := range array {
			if c < n {
				place++
			} else if c == n {
				ct++
			}
		}
		for ct > 0 {
			outArray[place] = n
			ct--
			place++
		}
	}
	return outArray
}

func ribbonNeeded(array []int) int {
	return (array[0] * array[1] * array[2]) + array[0] + array[0] + array[1] + array[1]
}

func paperNeeded(sides []int) int {
	total := 0
	minside := 0
	for i, s := range sides {
		for _, s2 := range sides[i+1:] {
			area := s * s2
			total = total + ((len(sides) - 1) * area)
			if area < minside || minside == 0 {
				minside = area
			}
		}
	}
	total += minside

	return total
}

func main() {
	fmt.Println(readFile(filePath))
	//fmt.Println("2x3x4: ", paperNeeded([]int{2, 3, 4}))
	//fmt.Println("1x1x10: ", paperNeeded([]int{1, 1, 10}))
	//infile := readFile(filePath)
	fmt.Println("28x28x29:", ribbonNeeded([]int{28, 28, 29}))
	fmt.Println("1x1x10:", ribbonNeeded([]int{1, 1, 10}))

	//fmt.Print(infile)
	//readFile(filePath)

}
