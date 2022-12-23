package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Parses input file for initial crate diagram and instructions - then moves crates according to the specified instructions w/ respect to the
model of CrateMover900x that is being used.
*/
func main() {
	part1Crates := moveCratesAndGetTopCrates(CrateMover9000Func)
	log.Printf("Part 1: These are the %v top crates after the move: %v", len(part1Crates), strings.Join(part1Crates, ""))
	part2Crates := moveCratesAndGetTopCrates(CrateMover9001Func)
	log.Printf("Part 2: These are the %v top crates after the move: %v", len(part2Crates), strings.Join(part2Crates, ""))

}

/*
Moves crates based on the provided inputfile and particular model of CrateMover900x
*/
func moveCratesAndGetTopCrates(crateMoverFunc CrateMoverFunction) []string {
	crateStack, moveInstructions := parseInputFile()
	for _, instructions := range moveInstructions {
		// Take the top N crates from the first stack, reverse them
		fromStack := crateMoverFunc(crateStack, instructions)

		// append crates from first stack to the second stack
		toStack := crateStack.CopyStack(instructions.To)
		newStack := append(fromStack, toStack...)
		crateStack[instructions.To] = newStack

		// Pop the first N crates off the first stack and reassign
		crateStack[instructions.From] = crateStack[instructions.From][instructions.Quantity:]
	}

	numStacks := len(crateStack)
	topCrates := []string{}
	for i := 1; i <= numStacks; i++ {
		topCrates = append(topCrates, string(crateStack[i][0]))
	}
	return topCrates
}

/*
Parses the provided input file for the initial state of the crate stack and the move instructions to apply on the crate stack
*/
func parseInputFile() (CrateStack, []MoveInstructions) {
	data, _ := os.ReadFile("resources/input")
	lines := strings.Split(string(data), "\n")
	crateStack, moveLinesStart := parseInitialCrateDiagram(lines)
	moveInstructions := parseMoveInstructions(lines[moveLinesStart:])
	return crateStack, moveInstructions
}

/*
Parses the initial state of the crate stack from the given input file
Returns the initialized CrateStack and line # where move set starts in input file
*/
func parseInitialCrateDiagram(lines []string) (CrateStack, int) {
	var initialCrateStack CrateStack = CrateStack{}
	var moveLinesStart int
	for i, line := range lines {
		for j := 0; j <= len(line); j += 4 {
			crate := line[j : j+3]
			crateId := rune(crate[1])

			/*
			 This is a hack to stop processing when we get to line w/ only column numbers
			*/
			_, err := strconv.Atoi(string(crateId))
			if err == nil {
				// Skip the current line and the following blank line to get to start of move instructions
				moveLinesStart = i + 2
				break
			}

			// Add the crateId rune to the relevant column if it is non-empty
			if crateId != ' ' {
				columnId := (j / 4) + 1
				columnCrates, _ := initialCrateStack[columnId]
				initialCrateStack[columnId] = append(columnCrates, crateId)
			}
		}
		if moveLinesStart != 0 {
			break
		}
	}

	return initialCrateStack, moveLinesStart
}

/*
Parses the set of instructions listed below the crate diagram, used to determine how many crates move from one stack to another
*/
func parseMoveInstructions(lines []string) []MoveInstructions {
	var instructions []MoveInstructions = []MoveInstructions{}
	for _, line := range lines {
		splitLine := strings.Split(line, " ")
		quantity, _ := strconv.Atoi(splitLine[1])
		from, _ := strconv.Atoi(splitLine[3])
		to, _ := strconv.Atoi(splitLine[5])
		instructions = append(instructions, MoveInstructions{quantity, from, to})
	}
	return instructions
}

/*
CrateMoverFunction describes the way in which a given model of 'CrateMover900x' moves crates given some CrateStack and MoveInstructions
*/
type CrateMoverFunction func(CrateStack, MoveInstructions) []rune

func CrateMover9000Func(crateStack CrateStack, instructions MoveInstructions) []rune {
	// Take the top N crates from the first stack, reverse them
	fromStack := crateStack.CopyStack(instructions.From)[:instructions.Quantity]
	for i, j := 0, len(fromStack)-1; i < j; i, j = i+1, j-1 {
		fromStack[i], fromStack[j] = fromStack[j], fromStack[i]
	}
	return fromStack
}

func CrateMover9001Func(crateStack CrateStack, instructions MoveInstructions) []rune {
	return crateStack.CopyStack(instructions.From)[:instructions.Quantity]
}

/*
CrateStack is a mapping of column ids to crate stacks labeled by some character
*/
type CrateStack map[int][]rune

func (cs CrateStack) CopyStack(id int) []rune {
	originalArr := cs[id]
	newArr := make([]rune, len(originalArr))
	copy(newArr, originalArr)
	return newArr
}

/*
MoveInstructions represents how many crates should be move, where they should be moved from, and where they should be moved to
*/
type MoveInstructions struct {
	Quantity int
	From     int
	To       int
}
