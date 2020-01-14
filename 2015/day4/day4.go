package main

import (
	"crypto/md5"
	"fmt"
	"io"
)

func getHash(s string) string {
	h := md5.New()
	_, err := io.WriteString(h, s)
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%x", h.Sum(nil)) //%x return hex version of hash
}

func mineAdventCoints(s string) string {
	ct := 0
	for {
		testKey := fmt.Sprintf("%s%d", s, ct)
		hashed := getHash(testKey)
		if hashed[:5] == "00000" {
			return fmt.Sprint(testKey, ":", hashed)
		}
		ct++
	}
}

func mineAdventCointsSix(s string) string {
	ct := 0
	for {
		testKey := fmt.Sprintf("%s%d", s, ct)
		if ct%1000000 == 0 {
			fmt.Printf("%d\r", ct)
		}
		if getHash(testKey)[:6] == "000000" {
			return fmt.Sprint(testKey, ":", getHash(testKey))
		}
		ct++
	}
}

func main() {
	i1 := "abcdef"
	i2 := "pqrstuv"
	i3 := "bgvyzdsv"

	//input := "bgvyzdsv"
	fmt.Println(mineAdventCoints(i1))
	fmt.Println(mineAdventCoints(i2))
	fmt.Println(mineAdventCoints(i3))
	fmt.Println(mineAdventCointsSix(i1))
	fmt.Println(mineAdventCointsSix(i2))
	fmt.Println(mineAdventCointsSix(i3))
	//getHash(input3
}
