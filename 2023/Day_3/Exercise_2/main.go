package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"path"
	"strconv"

	"golang.org/x/exp/slices"
)

const file = "input.txt"

var nrList = []string{
	"1", "2", "3", "4", "5", "6", "7", "8", "9", "0",
}

type line struct {
	previousLinePGears map[int]int // location
	previousLineNrs    map[int]int // location and nr
	currentLineNrs     map[int]int
	currentLinePGEARS  map[int]int
}

func (cLine *line) ExtractPartNumbers() int {
	sum := 0
	for _, p := range cLine.currentLineSymbols {
		if _, ok := cLine.currentLineNrs[p-1]; ok {
			fmt.Printf("valid nr found at %d\n", p-1)
			sum += checkAdjacentNr(p-1, cLine.currentLineNrs)

		}
		if _, ok := cLine.currentLineNrs[p+1]; ok {
			fmt.Printf("valid nr found at %d\n", p+1)
			sum += checkAdjacentNr(p+1, cLine.currentLineNrs)
		}
	}
	// symbols from previous line
	for _, p := range cLine.previousLineSymbols {
		if _, ok := cLine.currentLineNrs[p-1]; ok {
			fmt.Printf("valid nr found at %d\n", p-1)
			sum += checkAdjacentNr(p-1, cLine.currentLineNrs)
		}
		if _, ok := cLine.currentLineNrs[p]; ok {
			fmt.Printf("valid nr found at %d\n", p)
			sum += checkAdjacentNr(p, cLine.currentLineNrs)
		}
		if _, ok := cLine.currentLineNrs[p+1]; ok {
			fmt.Printf("valid nr found at %d\n", p+1)
			sum += checkAdjacentNr(p+1, cLine.currentLineNrs)
		}
	}
	// logic to check currentLineSymbols adjacent to previousLineNrs
	for _, p := range cLine.currentLineSymbols {
		if _, ok := cLine.previousLineNrs[p-1]; ok {
			fmt.Printf("TESTvalid nr found at %d\n", p-1)
			sum += checkAdjacentNr(p-1, cLine.previousLineNrs)
		}
		if _, ok := cLine.previousLineNrs[p]; ok {
			fmt.Printf("valid nr found at %d\n", p)
			sum += checkAdjacentNr(p, cLine.previousLineNrs)
		}
		if _, ok := cLine.previousLineNrs[p+1]; ok {
			fmt.Printf("valid nr found at %d\n", p+1)
			sum += checkAdjacentNr(p+1, cLine.previousLineNrs)
		}
	}
	return sum
}

func parseLine(l string) line {
	newLine := line{
		previousLineNrs: map[int]int{},
		currentLineNrs:  map[int]int{},
	}
	for i, c := range l {
		if string(c) != "." && !slices.Contains(nrList, string(c)) {
			newLine.currentLineSymbols = append(newLine.currentLineSymbols, i)
		}
		if slices.Contains(nrList, string(c)) {
			n, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal("failed to convert string to int")
			}
			newLine.currentLineNrs[i] = n
		}
	}
	return newLine
}

func main() {
	sum := 0
	previousLine := line{
		previousLineNrs: map[int]int{},
		currentLineNrs:  map[int]int{},
	}

	scanner := fileScanner()
	for scanner.Scan() {
		// get info for currentLine
		currentLine := parseLine(scanner.Text())
		currentLine.previousLineSymbols = previousLine.currentLineSymbols
		currentLine.previousLineNrs = previousLine.currentLineNrs

		sum += currentLine.ExtractPartNumbers()

		// update currentLine to previousLine
		previousLine.previousLineNrs = previousLine.currentLineNrs
		previousLine.previousLineSymbols = previousLine.currentLineSymbols
		previousLine.currentLineNrs = currentLine.currentLineNrs
		previousLine.currentLineSymbols = currentLine.currentLineSymbols
	}
	fmt.Println(sum)
}

func checkAdjacentNr(p int, m map[int]int) int {
	var number []int
	if a, ok1 := m[p-1]; ok1 {
		if b, ok2 := m[p-2]; ok2 {
			number = append(number, b)
			delete(m, p-2)
		}
		number = append(number, a)
		delete(m, p-1)
	}
	number = append(number, m[p])
	delete(m, p)
	if c, ok3 := m[p+1]; ok3 {
		number = append(number, c)
		delete(m, p+1)
		if d, ok4 := m[p+2]; ok4 {
			number = append(number, d)
			delete(m, p+2)
		}
	}
	singleNumber := 0
	for i, d := range number {
		singleNumber += powInt(10, len(number)-i-1) * d
	}
	fmt.Println(singleNumber)
	return singleNumber
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
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
