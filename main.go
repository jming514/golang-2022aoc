package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func bubbleSort(input []int) {
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input)-1; j++ {
			if input[j] > input[j+1] {
				temp := input[j]
				input[j] = input[j+1]
				input[j+1] = temp
			}
		}
	}
}

func main() {
	body, error := ioutil.ReadFile("day3.txt")

	if error != nil {
		fmt.Printf("Error reading file: %v", error)
	}

	// split every row into an element in a slice
	result := strings.Split(string(body), "\n")

	final := ""
	// get shared characters from each row
	for i := 0; i < len(result); i++ {
		firstHalf := result[i][0:(len(result[i]) / 2)]
		secondHalf := result[i][(len(result[i])/2)+1:]

		shared := ""
		for j := 0; j < len(firstHalf); j++ {
			for k := 0; k < len(secondHalf); k++ {
				if firstHalf[j] == secondHalf[k] && !strings.Contains(shared, string(firstHalf[j])) {
					shared += string(firstHalf[j])
				}
			}
		}

		final += shared
	}

	fmt.Printf("Final: %v\n", final)

	sum := 0
	for i := 0; i < len(final); i++ {

		if int(final[i]) < int('a') {
			sum += int(final[i] - 'A' + 1 + 26)
		} else {
			sum += int(final[i] - 'a' + 1)
		}
		// fmt.Printf("%v: %v\n", string(final[i]), int(final[i]-'a'+1))
	}
	fmt.Print(sum)
}

func dayTwo() {
	// 1 A = Rock         X = Rock/LOSE
	// 2 B = Paper        Y = Paper/DRAW
	// 3 C = Scissors     Z = Scissors/WIN
	// 1 = lose, 3 = draw, 6 = win
	body, error := ioutil.ReadFile("day2.txt")

	if error != nil {
		fmt.Printf("Error reading file: %v", error)
	}

	result := strings.Split(string(body), "\n")

	score := 0

	for i := 0; i < len(result); i++ {

		round := strings.Split(result[i], " ")
		// fmt.Printf("round %v: %v\n", i, round)

		if round[0] == "A" {
			if round[1] == "X" {
				score += 3 // 0 + 3
			} else if round[1] == "Y" {
				score += 4 // 3 + 1
			} else if round[1] == "Z" {
				score += 8 // 6 + 2
			}
		} else if round[0] == "B" {
			if round[1] == "X" {
				score += 1 // 0 + 1
			} else if round[1] == "Y" {
				score += 5 // 3 + 2
			} else if round[1] == "Z" {
				score += 9 // 6 + 3
			}
		} else if round[0] == "C" {
			if round[1] == "X" {
				score += 2 // 0 + 2
			} else if round[1] == "Y" {
				score += 6 // 3 + 3
			} else if round[1] == "Z" {
				score += 7 // 6 + 1
			}
		}
	}

	fmt.Printf("Score: %v", score)
}

func dayOne() {
	body, error := ioutil.ReadFile("day1.txt")

	if error != nil {
		fmt.Printf("Error reading file: %v", error)
	}

	result := strings.Split(string(body), "\n")

	current := 0
	var sums []int

	// n
	for i := 0; i < len(result); i++ {
		if result[i] == "" {
			sums = append(sums, current)
			current = 0
			continue
		}

		v, err := strconv.Atoi(result[i])

		if err != nil {
			panic(err)
		}

		current += v
	}

	// n^2
	bubbleSort(sums)

	finalSum := 0
	for _, a := range sums[(len(sums) - 3):] {
		finalSum += a
	}

	fmt.Printf("Last 3 elements of the array: %v\n", sums[(len(sums)-3):])
	fmt.Printf("Sum: %v", finalSum)
}
