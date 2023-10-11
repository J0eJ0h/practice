package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readInput() []string {
	//return []string{"30373", "25512", "65332", "33549", "35390"}

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

type Tree struct {
	Height, row, column int
	forest              *Forest
	Visible             bool
}

type Forest struct {
	rows, columns int
	trees         [][]*Tree //row, column
}

func readTrees(input []string) *Forest {
	f := &Forest{rows: len(input), columns: len(input[0])}

	f.trees = make([][]*Tree, f.rows, f.rows)
	for r := 0; r < f.rows; r++ {
		f.trees[r] = make([]*Tree, f.columns, f.columns)
		for c := 0; c < f.columns; c++ {
			height, _ := strconv.Atoi(string(input[r][c]))
			f.trees[r][c] = &Tree{Height: height, row: r, column: c, Visible: false, forest: f}
		}
	}
	return f
}

func (f *Forest) TreeHeight(r, c int) int {
	if r < 0 || c < 0 || r >= f.rows || c >= f.columns {
		return -1
	}
	return f.trees[r][c].Height
}

func (t *Tree) UpdateVis(height int) int {
	if t.Height > height {
		height = t.Height
		t.Visible = true
	}
	return height
}

func (f *Forest) TotalVisible() int {
	// do each direction
	// left to right
	for r := 0; r < f.rows; r++ {
		f.trees[r][0].Visible = true
		height := f.trees[r][0].Height
		for c := 1; c < f.columns; c++ {
			height = f.trees[r][c].UpdateVis(height)
		}
	}

	for c := 0; c < f.columns; c++ {
		f.trees[0][c].Visible = true
		height := f.trees[0][c].Height
		for r := 1; r < f.rows; r++ {
			height = f.trees[r][c].UpdateVis(height)
		}
	}

	for r := f.rows - 1; r >= 0; r-- {
		f.trees[r][f.columns-1].Visible = true
		height := f.trees[r][f.columns-1].Height
		for c := f.columns - 2; c >= 0; c-- {
			height = f.trees[r][c].UpdateVis(height)
		}
	}

	for c := f.columns - 1; c >= 0; c-- {
		f.trees[f.rows-1][c].Visible = true
		height := f.trees[f.rows-1][c].Height
		for r := f.rows - 2; r >= 0; r-- {
			height = f.trees[r][c].UpdateVis(height)
		}
	}

	// add em up
	total := 0
	for r := 0; r < f.rows; r++ {
		for c := 0; c < f.columns; c++ {

			p := " "
			if f.trees[r][c].Visible {
				total++
				p = fmt.Sprintf("%v", f.trees[r][c].Height)
			}
			fmt.Printf(p)
		}
		fmt.Println("")
	}

	return total
}

func part1() {
	input := readInput()
	fmt.Printf("part1 input length: %v\n", len(input))

	f := readTrees(input)

	fmt.Printf("Total visible: %v\n", f.TotalVisible())
}

func (f *Forest) ScoreTree(row, column int) (int, int, int, int, int) {
	h := f.TreeHeight(row, column)
	up, down, left, right := 0, 0, 0, 0
	for r := row - 1; r >= 0; r-- {
		left++
		if f.TreeHeight(r, column) >= h {
			break
		}
	}
	for r := row + 1; r < f.rows; r++ {
		right++
		if f.TreeHeight(r, column) >= h {
			break
		}
	}
	for c := column - 1; c >= 0; c-- {
		up++
		if f.TreeHeight(row, c) >= h {
			break
		}
	}
	for c := column + 1; c < f.columns; c++ {
		down++
		if f.TreeHeight(row, c) >= h {
			break
		}
	}
	return up * down * left * right, up, down, left, right
}

func part2() {
	input := readInput()
	fmt.Printf("part2 input length: %v\n", len(input))

	f := readTrees(input)

	score := 0
	for r := 0; r < f.rows; r++ {
		for c := 0; c < f.columns; c++ {
			s, _, _, _, _ := f.ScoreTree(r, c)
			if s > score {
				score = s
			}
		}
	}

	fmt.Printf("Best score: %v\n", score)
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
