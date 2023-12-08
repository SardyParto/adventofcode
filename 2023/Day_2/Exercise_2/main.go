package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
)

const (
	file = "input.txt"
)

func main() {
	sumPowers := 0
	scanner := fileScanner()
	for scanner.Scan() {
		line := scanner.Text()
		sumPowers += gameLimitPower(line)
	}

	fmt.Println(sumPowers)
}

func gameLimitPower(s string) int {
	var blueLimit int
	var greenLimit int
	var redLimit int
	reBlue, _ := regexp.Compile("[0-9]+ blue")
	reGreen, _ := regexp.Compile("[0-9]+ green")
	reRed, _ := regexp.Compile("[0-9]+ red")
	subsets := strings.Split(s, ";")
	for _, set := range subsets {
		blueCubes := convertToInt(reBlue.FindString(set))
		greenCubes := convertToInt(reGreen.FindString(set))
		redCubes := convertToInt(reRed.FindString(set))
		if blueCubes > blueLimit {
			blueLimit = blueCubes
		}
		if redCubes > redLimit {
			redLimit = redCubes
		}
		if greenCubes > greenLimit {
			greenLimit = greenCubes
		}

	}
	power := blueLimit * redLimit * greenLimit
	return power
}

func convertToInt(s string) int {
	if s == "" {
		return 0
	}

	b, err := strconv.Atoi(strings.TrimSpace(fmt.Sprint(s[0:2])))
	if err != nil {
		fmt.Println("error converting to int")
	}
	return b
}

func fileScanner() *bufio.Scanner {
	wd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting working directory")
	}

	f, err := os.Open(path.Join(wd, file))
	if err != nil {
		log.Fatal("Error opening input file")
	}

	scanner := bufio.NewScanner(f)
	return scanner
}
