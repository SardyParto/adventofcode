package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

const file = "input.txt"

var numbers = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine",
}

func main() {
	var sum int
	wd, err := os.Getwd()
	fmt.Println(wd)
	if err != nil {
		log.Fatal("Error getting working directory")
	}

	f, err := os.Open(path.Join(wd, file))
	if err != nil {
		log.Fatal("Error opening file")
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		sum += extractNumbers(line)
	}
	fmt.Println(sum)
}

func extractNumbers(s string) int {
	var fDigit int
	var lDigit int

	for c := range s {
		for i, number := range numbers {
			digit := i + 1
			if digit >= 10 {
				digit -= 9
			}
			if strings.HasPrefix(s[c:], number) {
				if fDigit == 0 {
					fDigit = digit
				}
				lDigit = digit
			}
		}
	}
	return fDigit*10 + lDigit
}
