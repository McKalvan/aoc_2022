package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	ADDX = "addx"
	NOOP = "noop"
)

func main() {
	signalStrengthMap := parseInputFile()
	summedSignalStrength := signalStrengthMap.SumSignalStrengths(20, 60, 100, 140, 180, 220)
	log.Printf("Part 1: Sum of signal strengths at 20th, 60th, 100th, 140th, 180th, and 220th cycles is %v", summedSignalStrength)
}

func parseInputFile() SignalStrengthMap {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(data), "\n")
	return parseProgram(lines)
}

func parseProgram(cmds []string) SignalStrengthMap {
	result := map[int]int{}
	var currentCycle int
	var x int = 1

	for _, cmd := range cmds {
		currentCycle++
		result[currentCycle] = x * currentCycle
		if strings.HasPrefix(cmd, ADDX) {
			vStr := strings.Split(cmd, " ")[1]
			v, _ := strconv.Atoi(vStr)
			currentCycle++
			result[currentCycle] = x * currentCycle
			x += v
		}
	}
	return result
}

type SignalStrengthMap map[int]int

func (ssMap SignalStrengthMap) SumSignalStrengths(cycles ...int) int {
	var result int
	for _, cycle := range cycles {
		result += ssMap[cycle]
	}
	return result
}
