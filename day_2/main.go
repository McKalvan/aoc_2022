package main

import (
	"log"
	"os"
	"strings"
)

const (
	// Input path containing data provided by AOC 2022 Day 2
	INPUT_PATH = "resources/input"

	// P1 Moves
	P1_ROCK    = "A"
	P1_PAPER   = "B"
	P1_SCISSOR = "C"

	// P2 Moves (part 1)
	P2_ROCK    = "X"
	P2_PAPER   = "Y"
	P2_SCISSOR = "Z"

	// P2 Strategies (part 2)
	P2_LOSE = "X"
	P2_DRAW = "Y"
	P2_WIN  = "Z"

	// Point values for each particular move
	ROCK_POINTS    = 1
	PAPER_POINTS   = 2
	SCISSOR_POINTS = 3

	// Points values for each particular outcome
	DRAW_POINTS = 3
	WIN_POINTS  = 6
)

var (
	// Used to translate P2 XYZ moves in part 1 to P1 ABC equivalent moves
	P2_TO_P1_MAP = map[string]string{
		P2_ROCK:    P1_ROCK,
		P2_PAPER:   P1_PAPER,
		P2_SCISSOR: P1_SCISSOR,
	}

	// Mapping between a particular move and the points awarded for playing that move
	MOVE_SCORE_MAP = map[string]int{
		P1_ROCK:    ROCK_POINTS,
		P1_PAPER:   PAPER_POINTS,
		P1_SCISSOR: SCISSOR_POINTS,
	}

	// Describes win conditions of RPS for each move
	RESULT_MAP = map[string]string{
		P1_PAPER:   P1_ROCK,
		P1_ROCK:    P1_SCISSOR,
		P1_SCISSOR: P1_PAPER,
	}
)

/*
	Parses the input file and determines the total point values awarded to each player using the strategy provided by the elf for each part:
		- Part 1: Translate P2 XYZ move to corresponding ABC move and play
		- Part 2: Translate P2 move based on the strategy provided to P2 and the move made by P1
*/
func main() {
	lines := parseInputFile()
	p1ScorePart1, p2ScorePart1 := calculatePoints(lines, getP2MoveMapping)
	log.Printf("Part 2 Final scores: P1: %v, P2: %v", p1ScorePart1, p2ScorePart1)
	p1ScorePart2, p2ScorePart2 := calculatePoints(lines, determineP2Move)
	log.Printf("Part 2 Final scores: P1: %v, P2: %v", p1ScorePart2, p2ScorePart2)
}

/*
	Parses the user-specific input file provided by https://adventofcode.com/2022/day/2/input
*/
func parseInputFile() []string {
	data, _ := os.ReadFile(INPUT_PATH)
	return strings.Split(string(data), "\n")
}

/*
	Calculates the final score of each player given their move (or strategy for p2 in part 2)
*/
func calculatePoints(lines []string, converter p2MoveConverter) (int, int) {
	var p1Score int
	var p2Score int

	for _, line := range lines {
		p1Move, p2Move := getPlayerMoves(line)
		p2MoveTranslated := converter(p1Move, p2Move)

		resultScore1, resultScore2 := calculateResultOfGame(p1Move, p2MoveTranslated)
		moveScore1, moveScore2 := MOVE_SCORE_MAP[p1Move], MOVE_SCORE_MAP[p2MoveTranslated]

		p1Score += moveScore1 + resultScore1
		p2Score += moveScore2 + resultScore2
	}
	return p1Score, p2Score
}

/*
	Parses given line to get moves for P1 and P2
*/
func getPlayerMoves(line string) (string, string) {
	splitMoves := strings.Split(line, " ")
	return splitMoves[0], splitMoves[1]
}

/*
	Calculates the final result of the match given each players' move  
*/
func calculateResultOfGame(p1Move string, p2Move string) (int, int) {
	if p1Move != p2Move {
		p1WinCondition := RESULT_MAP[p1Move]
		if p1WinCondition == p2Move {
			return WIN_POINTS, 0
		} else {
			return 0, WIN_POINTS
		}
	}
	return DRAW_POINTS, DRAW_POINTS
}

/* 
	Used to make calculatePoints function more generic to part 1 and part 2 
*/
type p2MoveConverter func(string, string) string

/*
	Translates moves for P2's move to equivalent move in P1's move set
*/
func getP2MoveMapping(_ string, p2Move string) string {
	p2MoveTranslated, ok := P2_TO_P1_MAP[p2Move]
	if !ok {
		panic("Invalid P2 move " + p2Move)
	}
	return p2MoveTranslated
}

/*
	The strategy given in part 2 dictates the following rules for player 2:
		- If P2 has an X, you need to lose
		- If P2 has a Y, you need to draw
		- IF P2 has a Z, you need to win
	This function determines what move P2 should play given P1's move and the strategy given to P2
*/
func determineP2Move(p1Move string, p2Strategy string) string {
	if p2Strategy != P2_DRAW {
		p1WinCondition := RESULT_MAP[p1Move]
		if p2Strategy == P2_LOSE {
			return p1WinCondition
		} else {
			// The move that beats p1 is the move that beats what p1 beats
			return RESULT_MAP[p1WinCondition]
		}
	}
	return p1Move
}
