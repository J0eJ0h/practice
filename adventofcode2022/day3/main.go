package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"golang.org/x/exp/maps"
)

func readInput() []string {
	f, err := os.Open(".\\input.txt")
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

func scoreChar(c rune) int {

	if 'a' <= c && c <= 'z' {
		return 1 + (int)(c-'a')
	} else if 'A' <= c && c <= 'Z' {
		return 27 + (int)(c-'A')
	} else {
		return -1
	}
}

func commonChars(a, b string) []rune {
	cc := make(map[rune]struct{}) // lets make a set!
	for _, u := range a {
		for _, v := range b {
			if u == v {
				cc[v] = struct{}{}
			}
		}
	}

	return maps.Keys(cc)
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	totz := 0
	for _, s := range input {
		l := len(s)
		a, b := s[:l/2], s[l/2:]
		totz += scoreChar(commonChars(a, b)[0])
	}
	fmt.Printf("Totz: %v\n", totz) // should be 8394  (PART 2 refactor prep)

}

func part2() {
	fmt.Println("part2")

	input := readInput()
	if len(input)%3 != 0 {
		log.Fatal("BAD INPUT. NO COOKIE")
	}

	totz := 0
	for i := 0; i < len(input); i += 3 {
		a, b, c := input[i], input[i+1], input[i+2]
		maybebadges := string(commonChars(a, b))
		badgers := commonChars(maybebadges, c)
		if len(badgers) != 1 {
			log.Fatalf("We don't need no stinking badgers: %v\n", len(badgers))
		}
		totz += scoreChar(badgers[0])
	}

	fmt.Printf("Totals: %v\n", totz)
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
