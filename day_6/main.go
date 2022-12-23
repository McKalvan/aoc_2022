package main

import (
	"log"
	"os"
)

/*
Part 1: Processes signal to identify start-of-packet marker (sequence of 4 distinct characters)
*/
func main() {
	inputSignal := parseInputFile()
	startOfPacketMarker := identifyStartOfMarker(inputSignal, 4)
	log.Printf("Part 1: For given signal, start-of-packet marker was detected after %v characters", startOfPacketMarker)
	startOfMessageMarker := identifyStartOfMarker(inputSignal, 14)
	log.Printf("Part 2: For given signal, start-of-message marker was detected after %v characters", startOfMessageMarker)

}

/*
Function to detect a marker based on N distinct/sequential characters in a signal
*/
func identifyStartOfMarker(signal string, distChars int) int {
	var startOfPacketMarker int
	empty := struct{}{}
	for i := 0; i <= len(signal)-distChars; i++ {
		if startOfPacketMarker == 0 {
			characterSet := make(map[rune]struct{}, distChars)
			characterGroup := signal[i : i+distChars]
			for j := 0; j < len(characterGroup); j++ {
				c := rune(characterGroup[j])
				characterSet[c] = empty
			}
			if len(characterSet) == distChars {
				// Add distChars to get the index of the last character in the start-of-marker
				startOfPacketMarker = i + distChars
				println(string(characterGroup))
			}
		}
	}
	return startOfPacketMarker
}

/*
Parses the user-specific input file provided by https://adventofcode.com/2022/day/6/input
*/
func parseInputFile() string {
	data, _ := os.ReadFile("resources/input")
	return string(data)
}
