package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput() []string {
	f, err := os.Open("./input.txt")
	if err != nil {
		log.Fatalf("Error reading input: %s\n", err)
		return nil
	}
	defer f.Close()

	var input []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		input = append(input, s.Text())
	}

	return input

}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error paring %v\n", err)
	}
	return i
}

func digit(r rune) int {
	i, err := strconv.Atoi(string(r))
	if err != nil {
		return -1
	}
	return i
}

func worddigit(s string) int {
	i, err := strconv.Atoi(string(s[0]))
	if err == nil {
		return i
	}
	if strings.HasPrefix(s, "one") {
		return 1
	}
	if strings.HasPrefix(s, "two") {
		return 2
	}
	if strings.HasPrefix(s, "three") {
		return 3
	}
	if strings.HasPrefix(s, "four") {
		return 4
	}
	if strings.HasPrefix(s, "five") {
		return 5
	}
	if strings.HasPrefix(s, "six") {
		return 6
	}
	if strings.HasPrefix(s, "seven") {
		return 7
	}
	if strings.HasPrefix(s, "eight") {
		return 8
	}
	if strings.HasPrefix(s, "nine") {
		return 9
	}
	return -1
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	total := 0
	for _, str := range input {
		runes := []rune(str)
		for _, r := range str {
			val := digit(r)
			if val >= 0 {
				fmt.Print(val)
				total += 10 * val
				break
			}
		}
		for i := len(runes) - 1; i >= 0; i-- {
			val := digit(runes[i])
			if val >= 0 {

				fmt.Println(val)
				total += val
				break
			}
		}

	}
	fmt.Println(total)
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	total := 0
	for _, str := range input {
		lastval := -1
		for i, _ := range str {
			val := worddigit(str[i:])
			if val >= 0 && lastval == -1 {
				fmt.Print(val)
				total += 10 * val
			}
			if val >= 0 {
				lastval = val
			}
		}
		fmt.Println(lastval)
		total += lastval

	}
	fmt.Println(total)
}

func main() {
	dopart := flag.Int("part", 2, "Specify question part")

	switch *dopart {
	case 1:
		part1()
	case 2:
		part2()
	default:
		fmt.Println("invalid part")
	}
}
