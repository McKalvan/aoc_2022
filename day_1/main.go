package main

import (
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	inputPath = "resources/input"
)

func main() {
	lines := parseInputFile()
	totalCalorieArr := getTotalCaloriesPerElf(lines)
	topElvesCalories := totalCalorieArr[0:3]
	log.Printf("Top 3 Elves: %v", topElvesCalories)

	var result int
	for _, calorie := range topElvesCalories {
		result += calorie
	}
	log.Printf("Total calories carried by top 3 elves: %v", result)
}

// Parses the user-specific input file provided by https://adventofcode.com/2022/day/1/input
func parseInputFile() []string {
	data, err := os.ReadFile(inputPath)
	check(err)
	return strings.Split(string(data), "\n")
}

// Determines the total calories carried by each elf (dictated by empty newline) and sorts in desc order
func getTotalCaloriesPerElf(lines []string) []int {
	var totalCalorieArr []int
	var currentElfInventory []int
	for _, line := range lines {
		if line == "" {
			totalCalories := GetTotalCalories(currentElfInventory)
			totalCalorieArr = append(totalCalorieArr, totalCalories)
			currentElfInventory = []int{}
		} else {
			calories, err := strconv.Atoi(line)
			check(err)
			currentElfInventory = append(currentElfInventory, calories)
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totalCalorieArr)))
	return totalCalorieArr
}

func GetTotalCalories(inv []int) int {
	var total int
	for _, inv := range inv {
		total += inv
	}
	return total
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
