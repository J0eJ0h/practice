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

func atoi64(s string) int64 {
	s = strings.TrimSpace(s)
	i, err := strconv.ParseInt(s, 10, 64)
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

type Line struct {
	x0, y0, dx, dy, det float64
}

func GetLine(x0, y0, dx, dy float64) *Line {
	return &Line{dx: dx, dy: dy, x0: x0, y0: y0, det: dx*y0 - dy*x0}
}

func (l *Line) getT(x float64) float64 {
	return (x - float64(l.x0)) / float64(l.dx)
}

func (l *Line) getY(x float64) float64 {
	return (x-float64(l.x0))*float64(l.dy)/float64(l.dx) + float64(l.y0)
}

func (l1 *Line) intersect(l2 *Line) float64 {
	xi := float64(l2.dx*l1.det - l1.dx*l2.det)
	xi /= float64(l2.dy*l1.dx - l1.dy*l2.dx)

	return xi
}

func ReadLines(input []string) []Line {
	lines := make([]Line, len(input))
	for i, s := range input {
		pv := strings.Split(s, "@")
		xyz0 := strings.Split(pv[0], ",")
		dxyz := strings.Split(pv[1], ",")
		lines[i] = *GetLine(
			float64(atoi64(xyz0[0])),
			float64(atoi64(xyz0[1])),
			float64(atoi64(dxyz[0])),
			float64(atoi64(dxyz[1])))

	}
	return lines
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))
	dopart1(input, 200000000000000, 400000000000000)
}
func dopart1(input []string, min, max float64) {
	total, ot := 0, 0

	lines := ReadLines(input)

	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			x := lines[i].intersect(&lines[j])
			y := lines[i].getY(x)

			if x >= min && x <= max {
				if y >= min && y <= max {
					ot += 1
					if lines[i].getT(x) >= 0 && lines[j].getT(x) >= 0 {
						fmt.Printf("%d,%d: %f,%f\n", i, j, x, y)
						total += 1
					} else {
						fmt.Printf("%d,%d\n", i, j)
					}
				} else {
					fmt.Println(y)
				}

			} else {
				fmt.Println(x)
			}

		}
	}

	fmt.Println(ot)
	fmt.Println(total)
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))
	total := 0

	fmt.Println(total)
}

func main() {
	dopart := flag.Int("part", 1, "Specify question part")

	switch *dopart {
	case 1:
		part1()
	case 2:
		part2()
	default:
		fmt.Println("invalid part")
	}
	l1 := GetLine(19, 13, -2, 1)
	fmt.Println(l1.getT(14.333))

	l2 := GetLine(18, 19, -1, -1)
	fmt.Println(l2.getT(14.333))
	fmt.Println(l1.intersect(l2))
	fmt.Println(l2.intersect(l1))

	l3 := GetLine(20, 25, -2, -2)
	x := l1.intersect(l3)
	fmt.Printf("%f, %f\n", x, l1.getY(x))

	fmt.Println(l2.intersect(l3))

	test :=
		`19, 13, 30 @ -2,  1, -2
18, 19, 22 @ -1, -1, -2
20, 25, 34 @ -2, -2, -4
12, 31, 28 @ -1, -2, -1
20, 19, 15 @  1, -5, -3`

	dopart1(strings.Split(test, "\n"), 7, 27)

	/*test = `467..114..
	...*......
	..35..633.
	......#...
	617*......
	.....+.58.
	..592.....
	......755.
	...$.*....
	.664.598..` */

	//dopart1(strings.Split(test, "\n"))
}
