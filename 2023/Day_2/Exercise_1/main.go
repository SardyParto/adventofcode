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
	file       = "input.txt"
	redLimit   = 12
	greenLimit = 13
	blueLimit  = 14
)

func main() {
	sumPossible := 0
	scanner := fileScanner()
	for scanner.Scan() {
		line := scanner.Text()
		possible, gameNr := gameIsPossible(line)
		fmt.Printf("possible: %v, gameNr: %d\n", possible, gameNr)
		if possible {
			sumPossible += gameNr
		}
	}
	fmt.Println(sumPossible)
}

func gameIsPossible(s string) (bool, int) {
	var gameNr int
	reBlue, _ := regexp.Compile("[0-9]+ blue")
	reGreen, _ := regexp.Compile("[0-9]+ green")
	reRed, _ := regexp.Compile("[0-9]+ red")
	fmt.Sscanf(s, "Game %d", &gameNr)
	subsets := strings.Split(s, ";")
	for _, set := range subsets {
		blueCubes := reBlue.FindString(set)
		greenCubes := reGreen.FindString(set)
		redCubes := reRed.FindString(set)

		if blueLimit < convertToInt(blueCubes) {
			fmt.Printf("bluecubes over limit %d\n", convertToInt(blueCubes))
			return false, gameNr
		}
		if greenLimit < convertToInt(greenCubes) {
			fmt.Printf("greencubes over limit %d\n", convertToInt(greenCubes))
			return false, gameNr
		}
		if redLimit < convertToInt(redCubes) {
			fmt.Printf("redcubes over limit %d\n", convertToInt(redCubes))
			return false, gameNr
		}

	}
	return true, gameNr
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
