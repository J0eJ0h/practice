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

type round struct {
	r, g, b int
}

type game struct {
	id     int
	rounds []round
}

func parseRound(r string) round {
	ret := round{0, 0, 0}
	cubes := strings.Split(r, ",")
	for _, c := range cubes {
		cp := strings.Split(strings.TrimSpace(c), " ")
		switch cp[1] {
		case "red":
			ret.r = atoi(cp[0])
		case "blue":
			ret.b = atoi(cp[0])
		case "green":
			ret.g = atoi(cp[0])
		default:
			fmt.Printf("WAT? %s\n", cp[1])
		}
	}

	return ret
}

func parseRounds(rs string) []round {
	rounds := make([]round, 0)
	for _, r := range strings.Split(rs, ";") {
		rounds = append(rounds, parseRound(r))
	}
	return rounds
}

func parseLine(l string) game {
	parsed := strings.Split(l, ":")
	id := atoi(strings.Split(parsed[0], " ")[1])

	g := game{id: id, rounds: parseRounds(parsed[1])}

	return g
}

func checkGame(g game) bool {
	for _, r := range g.rounds {
		if r.r > 12 || r.g > 13 || r.b > 14 {
			return false
		}
	}
	return true
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	total := 0
	for _, l := range input {
		g := parseLine(l)
		if checkGame(g) {
			total += g.id
		}
	}

	fmt.Println(total)
}

func gamePower(g game) int {
	mr, mg, mb := 0, 0, 0
	for _, r := range g.rounds {
		if r.r > mr {
			mr = r.r
		}
		if r.g > mg {
			mg = r.g
		}
		if r.b > mb {
			mb = r.b
		}
	}
	return mr * mg * mb
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	total := 0

	for _, l := range input {
		g := parseLine(l)
		total += gamePower(g)
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
