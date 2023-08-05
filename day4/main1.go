package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func splitter(r rune) bool {
	if r == ',' || r == '-' {
		return true
	}

	return false
}

func part1() {
	file, err := os.Open("input1.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {
		line := scanner.Text()

		results := strings.FieldsFunc(line, splitter)

		r0, _ := strconv.Atoi(results[0])
		r1, _ := strconv.Atoi(results[1])
		r2, _ := strconv.Atoi(results[2])
		r3, _ := strconv.Atoi(results[3])

		if r2 >= r0 && r1 >= r2 {
			count++
		} else if r0 >= r2 && r3 >= r0 {
			count++
		}
	}

	fmt.Printf("part1: %v\n", count)
}

func main() {
	// part1()
	part2()
}
