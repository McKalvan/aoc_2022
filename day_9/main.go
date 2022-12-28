package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	RIGHT = "R"
	LEFT  = "L"
	UP    = "U"
	DOWN  = "D"
)

func main() {
	instructions := openInputFile()
	log.Printf("Part 1: The tail of the rope visited %v unique positions", DetermineNumUniqueTailPositions(instructions, 2))
	// 6242 too high
	log.Printf("Part 2: The tail of the rope visited %v unique positions", DetermineNumUniqueTailPositions(instructions, 10))
}

func DetermineNumUniqueTailPositions(instructions []Instruction, numKnots int) int {
	initialKnots := make([]KnotPosition, numKnots)
	currentRope := RopePosition{numKnots, initialKnots}
	ropeHistory := []RopePosition{currentRope}
	for _, instruction := range instructions {
		currentRope = currentRope.MoveRope(instruction)
		ropeHistory = append(ropeHistory, currentRope)
	}

	var numUniquePositions int
	empty := struct{}{}
	tailHistory := map[KnotPosition]struct{}{}
	for _, rope := range ropeHistory {
		tail := rope.knotPositions[rope.numKnots-1]
		_, exists := tailHistory[tail]
		if !exists {
			numUniquePositions++
			tailHistory[tail] = empty
		}
	}
	return numUniquePositions
}

/*
Parse input file for AOC 2022 day_9 challenge
*/
func openInputFile() []Instruction {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}

	result := []Instruction{}
	splitLines := strings.Split(string(data), "\n")
	for _, line := range splitLines {
		splitLine := strings.Split(line, " ")
		moves, _ := strconv.Atoi(splitLine[1])
		/*
		 explode each move to its own step
		 this makes things a bit simpler but might be more costly performance-wise compared to range-based solutions
		*/
		for i := 0; i < moves; i++ {
			result = append(result, Instruction{splitLine[0], 1})
		}
	}
	return result
}

type Instruction struct {
	direction string
	moves     int
}

type RopePosition struct {
	numKnots      int
	knotPositions []KnotPosition
}

/*
Moves rope based on directions provided by input
Returns head and tail position after move
*/
func (rope RopePosition) MoveRope(instruction Instruction) RopePosition {
	// move head
	head := rope.knotPositions[0].MoveKnot(instruction)
	newPositions := []KnotPosition{head}

	// move all other knots
	for i := 1; i < rope.numKnots; i++ {
		tail := rope.knotPositions[i]
		tail = tail.MakeAdjacentTo(head)
		newPositions = append(newPositions, tail)
		head = tail
	}
	return RopePosition{rope.numKnots, newPositions}
}

type KnotPosition struct {
	x int
	y int
}

func (knot KnotPosition) MoveKnot(instruction Instruction) KnotPosition {
	var newKnot KnotPosition
	switch instruction.direction {
	case RIGHT:
		newKnot = KnotPosition{knot.x + instruction.moves, knot.y}
	case LEFT:
		newKnot = KnotPosition{knot.x - instruction.moves, knot.y}
	case UP:
		newKnot = KnotPosition{knot.x, knot.y + instruction.moves}
	case DOWN:
		newKnot = KnotPosition{knot.x, knot.y - instruction.moves}
	}
	return newKnot
}

func (knot KnotPosition) IsAdjacent(otherKnot KnotPosition) bool {
	return math.Abs(float64(knot.x-otherKnot.x)) <= 1 && math.Abs(float64(knot.y-otherKnot.y)) <= 1
}

func (knot KnotPosition) MakeAdjacentTo(otherKnot KnotPosition) KnotPosition {
	if knot.IsAdjacent(otherKnot) {
		return knot
	}
	dX := otherKnot.x - knot.x
	dY := otherKnot.y - knot.y

	if dX < 0 {
		dX = int(math.Max(float64(dX), -1))
	} else {
		dX = int(math.Min(float64(dX), 1))
	}

	if dY < 0 {
		dY = int(math.Max(float64(dY), -1))
	} else {
		dY = int(math.Min(float64(dY), 1))
	}

	return KnotPosition{knot.x + dX, knot.y + dY}
}
