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

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatalf("Error paring %v\n", err)
	}
	return i
}

type Pos struct {
	x, y int
}

func sign(n int) int {
	if n < 0 {
		return -1
	}
	return 1
}

func nextTail(h, t Pos) Pos {
	xm, ym := 0, 0
	if h.x == t.x {
		if h.y-1 > t.y {
			ym = 1
		} else if h.y+1 < t.y {
			ym = -1
		}
	} else if h.y == t.y {
		if h.x-1 > t.x {
			xm = 1
		} else if h.x+1 < t.x {
			xm = -1
		}
	} else {
		xd, yd := h.x-t.x, h.y-t.y
		if xd < -1 || xd > 1 || yd < -1 || yd > 1 {
			xm, ym = sign(xd), sign(yd)
		}
	}
	return Pos{t.x + xm, t.y + ym}
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	tailPoop := make(map[Pos]struct{})
	head, tail := Pos{0, 0}, Pos{0, 0}
	tailPoop[tail] = struct{}{}

	for _, v := range input {
		f := strings.Fields(v)
		count := atoi(f[1])
		for i := 0; i < count; i++ {
			switch f[0] {
			case "U":
				head.y++
			case "D":
				head.y--
			case "L":
				head.x--
			case "R":
				head.x++
			}
			tail = nextTail(head, tail)
			tailPoop[tail] = struct{}{}
			fmt.Printf("%v,%v\n", head, tail)
		}
	}

	fmt.Printf("Visited: %v\n", len(tailPoop))
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	tailPoop := make(map[Pos]struct{})
	knots := make([]Pos, 10, 10)
	for i := range knots {
		knots[i] = Pos{0, 0}
	}
	tailPoop[knots[0]] = struct{}{}

	for _, v := range input {
		f := strings.Fields(v)
		count := atoi(f[1])
		for i := 0; i < count; i++ {
			switch f[0] {
			case "U":
				knots[0].y++
			case "D":
				knots[0].y--
			case "L":
				knots[0].x--
			case "R":
				knots[0].x++
			}
			for i := 1; i < 10; i++ {
				knots[i] = nextTail(knots[i-1], knots[i])
			}
			tailPoop[knots[9]] = struct{}{}
		}
	}

	fmt.Printf("Visited: %v\n", len(tailPoop))
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
