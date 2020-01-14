package main

import (
	"bufio"
	"fmt"
	"os"
)

const vowels string = "aeiou"
const filePath = "./input.txt"

func checkFile(filePath string) int {
	inFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	reader := bufio.NewScanner(inFile)
	nice := 0
	count := 0
	for reader.Scan() {
		count++
		t := reader.Text()
		v := countVowels(t, vowels)
		d := doubleLetter(t)
		c := noNoStrings(t)
		//fmt.Printf("%q:, vowels: %d, double: %t, consec: %t\n", t, v, d, c)
		if v > 2 && d && !c {
			nice++
			fmt.Printf("%q: vowels: %d, double: %t,\t consec: %t\t Nice \t\t%d of %d\n", t, v, d, c, nice, count)
		} else {
			fmt.Printf("%q: vowels: %d, double: %t,\t consec: %t\t Naughty \t%d of %d\n", t, v, d, c, nice, count)
		}
	}
	return nice
}

func countVowels(input, vowels string) int {
	vowelCount := 0
	for _, r := range input {
		for _, v := range vowels {
			if r == v {
				vowelCount++
			}
		}
	}
	return vowelCount
}

func doubleLetter(input string) bool {
	var prev rune
	for _, r := range input {
		if r == prev {

			return true
		}
		prev = r
	}
	return false
}

func noNoStrings(input string) bool {
	runes := []rune(input)
	for i, r := range runes {
		if i+1 < len(runes) {
			if r == 97 || r == 99 || r == 112 || r == 120 {
				if runes[i+1] == r+1 {
					return true
				}

				//if input[i+1] == r+1 {
				//	fmt.Printf("prev: %q r: %q\n", prev, r)
				//	return true
			}
		}
	}

	return false
}

func consecLetters(input string) bool {
	var prev rune
	for _, r := range input {
		if r == prev+1 {
			fmt.Printf("prev: %q r: %q\n", prev, r)
			return true
		}
		prev = r
	}

	return false
}

func checkTwo(input string) bool {
	runes := []rune(input)
	var pairRepeat bool = false
	var splitRepeat bool = false
	for i, r := range runes {
		if i+1 < len(runes) {
			if !pairRepeat {
				for i2, _ := range runes {
					if i2 > i+1 && i2+1 < len(runes) {
						if r == runes[i2] && runes[i+1] == runes[i2+1] {
							fmt.Printf("%q%q,%d %d\n", r, runes[i+1], i, i2)
							pairRepeat = true
						}
					}
				}
			}
			if !splitRepeat {
				if i+2 < len(runes) {
					if r == runes[i+2] {
						fmt.Printf("%q%q%q\n", runes[i], runes[i+1], runes[i+2])
						splitRepeat = true
					}
				}
			}
		}
	}
	return pairRepeat && splitRepeat
}

func checkFile2(filePath string) int {
	inFile, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	reader := bufio.NewScanner(inFile)
	nice := 0
	count := 0
	for reader.Scan() {
		count++
		text := reader.Text()
		check := checkTwo(text)
		fmt.Println(text, check, count, nice)
		if check {
			nice++
		}
	}
	return nice
}

func main() {

	//fmt.Println(checkFile(filePath))
	fmt.Println(checkFile2(filePath))

	//fmt.Println("qjhvhtzxzqqjkmpb")
	//fmt.Println(checkTwo("qjhvhtzxzqqjkmpb"))
	//fmt.Println("xxyxx")
	//fmt.Println(checkTwo("xxyxx"))
	//fmt.Println("uurcxstgmygtbstg")
	//fmt.Println(checkTwo("uurcxstgmygtbstg"))
	//fmt.Println("ieodomkazucvgmuy")
	//fmt.Println(checkTwo("ieodomkazucvgmuy"))

}
