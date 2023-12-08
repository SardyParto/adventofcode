package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strconv"
	"unicode"
)

const file = "input.txt"

func main() {
	var sum int64
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory")
	}

	f, err := os.Open(path.Join(wd, file))
	if err != nil {
		log.Fatal("Error opening input file")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += combineDigits(line)
	}
	fmt.Println(sum)
}

func combineDigits(s string) int64 {
	var fDigit int64
	var lDigit int64

	for _, character := range s {
		var err error
		if unicode.IsNumber(character) {
			if fDigit == 0 {
				fDigit, err = strconv.ParseInt(string(character), 10, 0)
				if err != nil {
					log.Fatal("error converting character to string")
				}
			}
			lDigit, err = strconv.ParseInt(string(character), 10, 0)
			if err != nil {
				log.Fatal("error converting character to string")
			}
		}
	}
	return fDigit*10 + lDigit
}
