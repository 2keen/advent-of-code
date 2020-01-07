package main

import (
	"bufio"
	"fmt"
	"os"
)

const filePath = "./input.txt"

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

type house struct {
	x, y int
}

func visitHouses(s string) *map[house]int {
	visits := make(map[house]int)
	currentHouse := house{0, 0}
	for _, r := range s {
		switch r {
		case 62: //'>'
			currentHouse.x++
		case 60: //'<'
			currentHouse.x--
		case 94: //'^'
			currentHouse.y++
		case 118: //'v'
			currentHouse.y--
		default:
			fmt.Println("invalid rune")
		}
		visits[currentHouse] = visits[currentHouse] + 1
	}
	return &visits
}

func visitHousesRobo(s string) *map[house]int {
	visits := make(map[house]int)
	currentHouse := house{0, 0}
	currentHouseRobo := house{0, 0}
	for i, r := range s {
		switch r {
		case 62: //'>'
			if i%2 == 0 {
				currentHouse.x++
			} else {
				currentHouseRobo.x++
			}
		case 60: //'<'
			if i%2 == 0 {
				currentHouse.x--
			} else {
				currentHouseRobo.x--
			}
		case 94: //'^'
			if i%2 == 0 {
				currentHouse.y++
			} else {
				currentHouseRobo.y++
			}
		case 118: //'v'
			if i%2 == 0 {
				currentHouse.y--
			} else {
				currentHouseRobo.y--
			}
		default:
			fmt.Println("invalid rune")
		}
		if i%2 == 0 {
			visits[currentHouse] = visits[currentHouse] + 1
		} else {
			visits[currentHouseRobo] = visits[currentHouseRobo] + 1
		}
	}
	return &visits
}

func main() {
	flightPath := readFile(filePath)
	houseVisits := visitHouses(flightPath)
	houseVisitsRobo := visitHousesRobo(flightPath)
	fmt.Println("houses visited:", len(*houseVisits))
	fmt.Println("  robo visited:", len(*houseVisitsRobo))
}
