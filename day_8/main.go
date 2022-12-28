package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	trees := parseInputFile()
	trees.AssignVisibility()
	log.Printf("Part 1: There are %v visible trees in the forest", trees.CountVisible())
	log.Printf("Part 2: The highest possible scenic score is %v", trees.GetMaxScenicScore())
}

/*
Parses input file for day_8 AOC 2022 task
*/
func parseInputFile() Forest {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}

	trees := [][]*Tree{}
	splitLines := strings.Split(string(data), "\n")
	for _, line := range splitLines {
		row := []*Tree{}
		for _, val := range line {
			heightVal, _ := strconv.Atoi(string(val))
			row = append(row, &Tree{heightVal, false})
		}
		trees = append(trees, row)
	}
	return trees
}

/*
Assigns visibility value to each tree
A tree is visible if it is taller than any tree that came before it in any given direction
*/
func (forest Forest) AssignVisibility() {

	// row visibility
	for _, row := range forest {
		// left to right visibility
		maxObserved := -1
		for _, tree := range row {
			if tree.height > maxObserved {
				tree.visible = true
				maxObserved = tree.height
			}
		}

		// right to left visibility
		maxObserved = -1
		for i := len(row) - 1; i >= 0; i-- {
			tree := row[i]
			if tree.height > maxObserved {
				tree.visible = true
				maxObserved = tree.height
			}
		}
	}

	// column visibility
	numCols := len(forest[0])
	for i := 0; i < numCols; i++ {
		maxObserved := -1
		for j := 0; j < len(forest); j++ {
			tree := forest[j][i]
			if tree.height > maxObserved {
				tree.visible = true
				maxObserved = tree.height
			}
		}

		maxObserved = -1
		for j := len(forest) - 1; j >= 0; j-- {
			tree := forest[j][i]
			if tree.height > maxObserved {
				tree.visible = true
				maxObserved = tree.height
			}
		}
	}
}

/*
Returns # of trees w/ visible field equal to true
*/
func (forest Forest) CountVisible() int {
	var visible int
	for _, row := range forest {
		for _, tree := range row {
			if tree.visible {
				visible++
			}
		}
	}
	return visible
}

/*
Get the value of the tree w/ the highest scenic score
scenic score is determined by multiplying the # of trees that can be scene from a given tree from each direction
*/
func (forest Forest) GetMaxScenicScore() int {
	maxRows := len(forest)
	maxCols := len(forest[0])
	var maxScenicScore int
	for i, row := range forest {
		for j, tree := range row {

			// visibility score to the right
			var rightScore int
			for z := j + 1; z < maxRows; z++ {
				otherTree := row[z]
				if otherTree.height < tree.height {
					rightScore++
				} else {
					rightScore++
					break
				}
			}

			// visibility score to the left
			var leftScore int
			for z := j - 1; z >= 0; z-- {
				otherTree := row[z]
				if otherTree.height < tree.height {
					leftScore++
				} else {
					leftScore++
					break
				}
			}

			// visibility score going down
			var downScore int
			for z := i + 1; z < maxCols; z++ {
				otherTree := forest[z][j]
				if otherTree.height < tree.height {
					downScore++
				} else {
					downScore++
					break
				}
			}

			// visibility score going up
			var upScore int
			for z := i - 1; z >= 0; z-- {
				otherTree := forest[z][j]
				if otherTree.height < tree.height {
					upScore++
				} else {
					upScore++
					break
				}
			}

			scenicScore := 1
			for _, score := range []int{leftScore, rightScore, downScore, upScore} {
				if score > 0 {
					scenicScore *= score
				}
			}

			if scenicScore > maxScenicScore {
				maxScenicScore = scenicScore
			}
		}
	}
	return maxScenicScore
}

/*
Tree represents an individual character in the input file
*/
type Tree struct {
	height  int
	visible bool
}

/*
Forest is an arrangement of trees in rows/columns
*/
type Forest [][]*Tree

/*
Generic function type used to align trees to assign them visibility values
*/
type AlignTreesFunc func(forest Forest) Forest

/*
Trees are already aligned left-to-right row-wise so this is a pass-through function
*/
func LeftAlign(forest Forest) Forest {
	return forest
}
