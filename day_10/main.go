package main

import (
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const (
	ADDX       = "addx"
	NOOP       = "noop"
	LIT_PIXEL  = "#"
	DARK_PIXEL = "."
)

func main() {
	signalStrengthMap := parseInputFile()
	summedSignalStrength := signalStrengthMap.SumSignalStrengths(20, 60, 100, 140, 180, 220)
	log.Printf("Part 1: Sum of signal strengths at 20th, 60th, 100th, 140th, 180th, and 220th cycles is %v", summedSignalStrength)
	log.Println("Part 2: Output of signal on CRT:")
	PrintCRTOutput(40, signalStrengthMap)
}

func parseInputFile() CycleRegisterMap {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return parseProgram(lines)
}

/*
Parses set of commands from input to CycleRegisterMap
Commands in input are limited to addx and noop
*/
func parseProgram(cmds []string) CycleRegisterMap {
	result := map[int]int{}
	var currentCycle int
	xRegister := 1

	for _, cmd := range cmds {
		// this handles noop and first cycle of addx
		currentCycle++
		result[currentCycle] = xRegister

		if strings.HasPrefix(cmd, ADDX) {
			// handles second cycle of addx
			vStr := strings.Split(cmd, " ")[1]
			v, _ := strconv.Atoi(vStr)

			currentCycle++
			result[currentCycle] = xRegister
			xRegister += v
		}
	}
	return result
}

type CycleRegisterMap map[int]int

/*
Sums the signal strengths (cycle * xRegister corresponding to cycle) of a given set of cycles
*/
func (cycleMap CycleRegisterMap) SumSignalStrengths(cycles ...int) int {
	var result int
	for _, cycle := range cycles {
		result += cycleMap[cycle] * cycle
	}
	return result
}

/*
Prints a sequence of capital letters based on info in cycleMap
*/
func PrintCRTOutput(cyclesPerRow int, cycleMap CycleRegisterMap) {
	for cycle := 1; cycle <= len(cycleMap); cycle++ {
		// update the currentWritePosition based on the current cycle and # of cycles per row
		currentWritePosition := (cycle - 1) % cyclesPerRow

		// update position of sprite based on xRegister val for current cycle
		spritePosition := cycleMap[cycle]

		// draw pixel in current position if visible in current cycle, EX w/in +/-1 of position being written
		if math.Abs(float64(spritePosition-currentWritePosition)) <= 1 {
			print(LIT_PIXEL)
		} else {
			print(DARK_PIXEL)
		}

		// check if new row needs to be written based on cycle val
		if cycle%cyclesPerRow == 0 {
			println()
		}
	}
}
