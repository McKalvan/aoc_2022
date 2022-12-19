package main

import (
	"log"
	"os"
	"strings"
	"unicode"
)

var priorityMap map[rune]int = map[rune]int{}

func main() {
	initPriorityMap()
	lines := openInputFile()
	evaluatePart1(lines)
	evaluatePart2(lines)
}

func initPriorityMap() {
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		priorityMap[r] = i
		priorityMap[unicode.ToUpper(r)] = i + 26
		i++
	}
}

func evaluatePart1(lines []string) {
	var prioritySum int
	for _, line := range lines {
		compartment1, compartment2 := splitComparments(line)
		misplacedItem := getMisplacedItem(compartment1, compartment2)
		prioritySum += priorityMap[misplacedItem]
	}
	log.Printf("Part 1: The sum of priorities for all misplaced items is %v", prioritySum)
}

func openInputFile() []string {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func splitComparments(ruckSackStr string) (string, string) {
	ruckSackSize := len(ruckSackStr)
	midPoint := ruckSackSize / 2
	compartment1 := ruckSackStr[0:midPoint]
	compartment2 := ruckSackStr[midPoint:ruckSackSize]
	return compartment1, compartment2
}

func getMisplacedItem(compartment1 string, compartment2 string) rune {
	// create empty bitmap
	itemBitMap := map[rune]bool{}

	// load the bitmap w/ items included in compartment 1
	for _, item := range compartment1 {
		itemBitMap[item] = true
	}

	var misplacedItem rune
	// check if compartment2 has any of the existing items added to bitmap from compartment 1
	for _, item := range compartment2 {
		_, exists := itemBitMap[item]
		if exists {
			misplacedItem = item
			break
		}
	}

	if misplacedItem == 0 {
		panic("There should be exactly one misplaced item in each rucksack")
	}

	return misplacedItem
}

func evaluatePart2(lines []string) {
	numGroupings := len(lines) / 3

	for i := 0; i <= numGroupings; i = i + 3 {
		group := lines[i : i+3]
		log.Printf("%v,%v,%v", group[0], group[1], group[2])
	}
}

func evaluateGroupTag(group [3]string) rune {

	// The groups item must be an item which is included at least once w/in every member of the groups rucksack
	sharedItems := map[rune]int{}
	for _, elfRucksack := range group {
		for _, item := range elfRucksack {	
			share
		} 
	}
}
