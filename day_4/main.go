package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := openInputFile()
	numSubsets := evaluatePart(lines, EvaluateSubset)
	log.Printf("Part 1: There are %v elf pairings in which the range that one elf is cleaning contains the entire range that the paired elf is cleaning", numSubsets)
	numIntersections := evaluatePart(lines, EvaluateIntersection)
	log.Printf("Part 2: There are %v elf pairings that have intersection assigments", numIntersections)
}

func openInputFile() []string {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func evaluatePart(lines []string, overlapFunc OverlapFunc) int {
	var numOverlaps int
	for _, elfPairings := range lines {
		assignment1, assignment2 := ParseElfAssignments(elfPairings)
		if overlapFunc(assignment1, assignment2) {
			numOverlaps++
		}
	}
	return numOverlaps
}

func ParseElfAssignments(elfPairings string) (ElfAssignment, ElfAssignment) {
	assignments := strings.Split(elfPairings, ",")
	return ElfAssignment{assignments[0]}, ElfAssignment{assignments[1]}
}

type OverlapFunc func(ElfAssignment, ElfAssignment) bool

func EvaluateSubset(ea1, ea2 ElfAssignment) bool {
	return ea1.IsSubsetOfPairedElf(ea2) || ea2.IsSubsetOfPairedElf(ea1)
}

func EvaluateIntersection(ea1, ea2 ElfAssignment) bool {
	return ea1.Intersects(ea2) || ea2.Intersects(ea1)
}

type ElfAssignment struct {
	sections string
}

func (ea ElfAssignment) GetRange() (int, int) {
	sectionRange := strings.Split(ea.sections, "-")
	low, _ := strconv.Atoi(sectionRange[0])
	high, _ := strconv.Atoi(sectionRange[1])
	return low, high
}

func (ea ElfAssignment) IsSubsetOfPairedElf(other ElfAssignment) bool {
	thisLow, thisHigh := ea.GetRange()
	otherLow, otherHigh := other.GetRange()
	return thisLow >= otherLow && thisHigh <= otherHigh
}

func (ea ElfAssignment) Intersects(other ElfAssignment) bool {
	thisLow, thisHigh := ea.GetRange()
	otherLow, otherHigh := other.GetRange()
	return (thisLow <= otherLow && thisHigh >= otherLow) || (thisHigh <= otherHigh && thisHigh >= otherLow)
}