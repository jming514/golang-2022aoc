package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func part2() {
	file, err := os.Open("input2.txt")
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
		} else if r2 <= r0 && r0 <= r3 {
			count++
		}
	}

	fmt.Printf("part2: %v\n", count)
}
