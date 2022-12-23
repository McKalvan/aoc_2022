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

func openInputFile() []string {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func evaluatePart1(lines []string) {
	var prioritySum int
	for _, line := range lines {
		compartment1, compartment2 := SplitCompartments(line)
		itemSet1, itemSet2 := StringToSet(compartment1), StringToSet(compartment2)
		misplacedItem := itemSet1.Intersection(itemSet2).GetItems()[0]
		prioritySum += priorityMap[misplacedItem]
	}
	log.Printf("Part 1: The sum of priorities for all misplaced items is %v", prioritySum)
}

func evaluatePart2(lines []string) {
	var prioritySum int
	numGroupings := len(lines) / 3
	for i := 0; i < numGroupings; i++ {
		group := lines[i*3 : 3*(i+1)]
		groupBadge := evaluateGroupBadge(group)
		prioritySum += priorityMap[groupBadge]
		println(i)
	}
	log.Printf("Total sum of priorities of each groups tag is %v", prioritySum)
}

func evaluateGroupBadge(group []string) rune {
	// The groups item must be an item which is included at least once w/in every member of the groups rucksack
	itemSet1, itemSet2, itemSet3 := StringToSet(group[0]), StringToSet(group[1]), StringToSet(group[2])
	intersection := itemSet1.Intersection(itemSet2).Intersection(itemSet3)
	return intersection.GetItems()[0]
}

func SplitCompartments(ruckSackStr string) (string, string) {
	ruckSackSize := len(ruckSackStr)
	midPoint := ruckSackSize / 2
	compartment1 := ruckSackStr[0:midPoint]
	compartment2 := ruckSackStr[midPoint:ruckSackSize]
	return compartment1, compartment2
}

func StringToSet(rs string) ItemSet {
	empty := struct{}{}
	result := make(map[rune]struct{})
	for _, item := range rs {
		result[item] = empty
	}
	return ItemSet{result}
}

type ItemSet struct {
	itemSet map[rune]struct{}
}

func (is ItemSet) GetItems() []rune {
	keys := make([]rune, 0, len(is.itemSet))
	for k := range is.itemSet {
		keys = append(keys, k)
	}
	return keys
}

func (is ItemSet) Equals(other ItemSet) bool {
	for item := range is.itemSet {
		_, exists := other.itemSet[item]
		if !exists {
			return false
		}
	}
	return true
}

func (is ItemSet) Intersection(other ItemSet) ItemSet {
	existsVal := struct{}{}
	intersection := map[rune]struct{}{}
	for item := range is.itemSet {
		_, exists := other.itemSet[item]
		if exists {
			intersection[item] = existsVal
		}
	}
	return ItemSet{intersection}
}
