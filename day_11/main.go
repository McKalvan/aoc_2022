package main

import (
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

const (
	NUM_LINES_MONKEY_DEFINITION = 7
	MONKEY                      = "Monkey "
	STARTING                    = "Starting items: "
	OPERATION                   = "Operation: new = "
	TEST                        = "Test: divisible by "
	IF_TRUE                     = "If true: throw to monkey "
	IF_FALSE                    = "If false: throw to monkey "
)

// Boo - global variables that are only relevant to one type
var worryStrategy WorryStrategy
var postInspectionWorryDivisor int64

func main() {
	worryStrategy = Part1WorryStrategy
	postInspectionWorryDivisor = 3
	totalMonkeyBusinessPart1 := CalculateTotalMonkeyBusiness(20, 2)
	log.Printf("Part 1: Total monkey business after %v rounds is %v", 20, totalMonkeyBusinessPart1)

	worryStrategy = Part2WorryStrategy
	postInspectionWorryDivisor = CalculateGlobalMod()
	totalMonkeyBusinessPart2 := CalculateTotalMonkeyBusiness(10000, 2)
	log.Printf("Part 1: Total monkey business after %v rounds is %v", 10000, totalMonkeyBusinessPart2)
}

func CalculateTotalMonkeyBusiness(numRounds int, numTopMonkeys int) int {
	monkeyMap := parseInput()
	// Run numRounds of monkey in the middle
	for i := 0; i < numRounds; i++ {
		for j := 0; j < len(monkeyMap); j++ {
			monkey := monkeyMap[j]
			monkey.TakeTurn(monkeyMap)
		}
	}

	monkeys := make([]*Monkey, len(monkeyMap))
	for i, monkey := range monkeyMap {
		monkeys[i] = monkey
	}
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].numberItemsInspected > monkeys[j].numberItemsInspected
	})

	monkeyBusiness := 1
	for _, monkey := range monkeys[:numTopMonkeys] {
		monkeyBusiness *= monkey.numberItemsInspected
	}
	return monkeyBusiness
}

func CalculateGlobalMod() int64 {
	monkeyMap := parseInput()
	var result int64 = 1
	for _, monkey := range monkeyMap {
		result *= int64(monkey.divisor)
	}
	return result
}

func parseInput() map[int]*Monkey {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	return parseMonkeys(strings.Split(string(data), "\n"))
}

/*
Parses input to mapping of monkey id to Monkey
The definition of each monkey takes up 7 lines, so there are len(lines)/7 monkeys
*/
func parseMonkeys(lines []string) map[int]*Monkey {
	monkeyMap := map[int]*Monkey{}
	for i := 0; i < len(lines)/NUM_LINES_MONKEY_DEFINITION; i++ {
		lineStart := i * NUM_LINES_MONKEY_DEFINITION
		monkeyDefinitionLines := lines[lineStart : lineStart+NUM_LINES_MONKEY_DEFINITION]
		monkey := LinesToMonkey(monkeyDefinitionLines)
		monkeyMap[monkey.id] = monkey
	}
	return monkeyMap
}

/*
Parses a group of 7 lines to a Monkey
Each line corresponds to part of the definition of a monkey
*/
func LinesToMonkey(lines []string) *Monkey {
	// monkey id
	idStr := strings.Split(lines[0], MONKEY)[1][0]
	id, _ := strconv.Atoi(string(idStr))
	// monkey starting items
	startingItems := ParseItems(lines[1])

	// operator
	operationStr := strings.Split(lines[2], OPERATION)[1]
	operation := ParseOperation(operationStr)

	// monkey divisor
	divisorStr := strings.Split(lines[3], TEST)[1]
	divisor, _ := strconv.Atoi(string(divisorStr))
	// true monkey id
	trueMonkeyStr := strings.Split(lines[4], IF_TRUE)[1]
	trueMonkeyId, _ := strconv.Atoi(string(trueMonkeyStr))
	// false monkey id
	falseMonkeyStr := strings.Split(lines[5], IF_FALSE)[1]
	falseMonkeyId, _ := strconv.Atoi(string(falseMonkeyStr))

	return &Monkey{
		id:            id,
		items:         startingItems,
		operator:      operation,
		divisor:       int64(divisor),
		trueMonkeyId:  trueMonkeyId,
		falseMonkeyId: falseMonkeyId,
	}
}

func ParseItems(itemsLine string) []int64 {
	startingItemsStr := strings.Split(itemsLine, STARTING)[1]
	startingItems := strings.Split(startingItemsStr, ", ")
	result := make([]int64, len(startingItems))
	for i, item := range startingItems {
		itemInt, _ := strconv.Atoi(item)
		result[i] = int64(itemInt)
	}
	return result
}

func ParseOperation(operationStr string) MonkeyBusiness {
	operationArr := strings.Split(operationStr, " ")
	operator := operationArr[1]
	rhsStr := operationArr[2]
	rhsVal, rhsIsStr := strconv.Atoi(rhsStr)
	rhsVal64 := int64(rhsVal)

	monkeyBusiness := func(old int64) int64 {
		if rhsIsStr != nil {
			rhsVal64 = old
		}

		var result int64
		switch operator {
		case "*":
			result = old * rhsVal64
		case "+":
			result = old + rhsVal64
		case "-":
			result = old - rhsVal64
		case "/":
			result = old / rhsVal64
		}
		return result
	}
	return monkeyBusiness
}

type MonkeyBusiness func(old int64) int64

type Monkey struct {
	id                   int
	items                []int64
	operator             MonkeyBusiness
	divisor              int64
	trueMonkeyId         int
	falseMonkeyId        int
	numberItemsInspected int
}

func (monkey *Monkey) TakeTurn(monkeyMap map[int]*Monkey) {
	for _, item := range monkey.items {
		// inspect the item
		targetMonkeyId, itemWorryLevel := monkey.InspectItem(item)
		monkey.numberItemsInspected++

		// give target monkey the item
		targetMonkey := monkeyMap[targetMonkeyId]
		targetMonkey.ReceiveItem(itemWorryLevel)
	}
	// clear inventory of monkey
	monkey.items = []int64{}
}

func (monkey *Monkey) InspectItem(item int64) (int, int64) {
	itemNewWorryLevel := worryStrategy(monkey.operator(item))

	var targetMonkeyId int
	if itemNewWorryLevel%monkey.divisor == 0 {
		targetMonkeyId = monkey.trueMonkeyId
	} else {
		targetMonkeyId = monkey.falseMonkeyId
	}
	return targetMonkeyId, itemNewWorryLevel
}

func (monkey *Monkey) ReceiveItem(item int64) {
	monkey.items = append(monkey.items, item)
}

type WorryStrategy func(int64) int64

func Part1WorryStrategy(item int64) int64 {
	return int64(math.Floor(float64(item / postInspectionWorryDivisor)))
}

func Part2WorryStrategy(item int64) int64 {
	return item % postInspectionWorryDivisor
}
