package main

import (
	"log"
	"math"
	"os"
	"strings"
)

const (
	START = "S"
	END   = "E"
)

func main() {
	lines := parseInput()
	topographyGraph := BuildTopographyGraph(lines)
	topographyGraph.AddEdges()

	shortestDistance := topographyGraph.FindShortestDistanceToTarget()
	log.Printf("Part 1: The shortest distance to the target is %v", shortestDistance)

	shortestDistanceFromLowestPoint := topographyGraph.FindShortestDistanceFromLowestToTarget()
	log.Printf("Part 2: The shortest distance to the target from an the lowest ('a') square is %v", shortestDistanceFromLowestPoint)
}

func parseInput() []string {
	data, err := os.ReadFile("resources/input")
	if err != nil {
		panic(err)
	}
	return strings.Split(string(data), "\n")
}

func BuildTopographyGraph(lines []string) TopographyGraph {
	topGraph := TopographyGraph{map[Coordinates]*MapNode{}, nil, nil}
	for i, line := range lines {
		for j, nodeVal := range line {
			mapNode := &MapNode{int(nodeVal), false, nil, nil}
			switch string(nodeVal) {
			case START:
				mapNode.Height = int('a')
				mapNode.IsExplored = true
				topGraph.StartingNode = mapNode

			case END:
				mapNode.Height = int('z')
				topGraph.TargetNode = mapNode
			}
			topGraph.AddNode(mapNode, Coordinates{j, i})
		}
	}
	return topGraph
}

type TopographyGraph struct {
	NodeMap      map[Coordinates]*MapNode
	StartingNode *MapNode
	TargetNode   *MapNode
}

func (topGraph *TopographyGraph) AddNode(node *MapNode, key Coordinates) {
	topGraph.NodeMap[key] = node
}

func (topGraph *TopographyGraph) AddEdges() {
	for nCoord, node := range topGraph.NodeMap {
		upNode, uExists := topGraph.NodeMap[nCoord.GetAdjacentSpace(0, 1)]
		if uExists && node.IsAccessible(upNode) {
			node.AppendAccessibleNode(upNode)
		}
		downNode, dExists := topGraph.NodeMap[nCoord.GetAdjacentSpace(0, -1)]
		if dExists && node.IsAccessible(downNode) {
			node.AppendAccessibleNode(downNode)
		}
		rightNode, rExists := topGraph.NodeMap[nCoord.GetAdjacentSpace(1, 0)]
		if rExists && node.IsAccessible(rightNode) {
			node.AppendAccessibleNode(rightNode)
		}
		leftNode, lExists := topGraph.NodeMap[nCoord.GetAdjacentSpace(-1, 0)]
		if lExists && node.IsAccessible(leftNode) {
			node.AppendAccessibleNode(leftNode)
		}
	}
}

// filter out just a's, run BFS on each
func (topGraph *TopographyGraph) FindShortestDistanceFromLowestToTarget() int {
	aNodes := []*MapNode{}
	for _, node := range topGraph.NodeMap {
		if rune(node.Height) == 'a' {
			aNodes = append(aNodes, node)
		}
	}
	minSteps := math.MaxInt
	for _, aNode := range aNodes {
		topGraph.ResetExploredNodes()

		aNode.IsExplored = true
		topGraph.StartingNode = aNode

		stepsToTarget := topGraph.FindShortestDistanceToTarget()
		if stepsToTarget > 0 && stepsToTarget < minSteps {
			minSteps = stepsToTarget
		}
	}
	return minSteps
}

func (topGraph *TopographyGraph) FindShortestDistanceToTarget() int {
	node := topGraph.GetShortestPathToTarget()
	var numSteps int
	for {
		if node.ParentNode == nil {
			break
		}
		numSteps++
		node = node.ParentNode
	}
	return numSteps
}

func (topGraph *TopographyGraph) GetShortestPathToTarget() *MapNode {
	stack := []*MapNode{topGraph.StartingNode}
	var head *MapNode
	for {
		head, stack = stack[0], stack[1:]
		if head == topGraph.TargetNode {
			break
		}
		for _, adjacentNode := range head.AccessibleNodes {
			if !adjacentNode.IsExplored {
				adjacentNode.IsExplored = true
				adjacentNode.ParentNode = head
				stack = append(stack, adjacentNode)
			}
		}
		if len(stack) == 0 {
			head = &MapNode{}
			break
		}
	}
	return head
}

func (topMap *TopographyGraph) PrintNumAccessibleNodes() {
	for i := 0; i < 40; i++ {
		for j := 0; j < 100; j++ {
			coordinate := Coordinates{j, i}
			node := topMap.NodeMap[coordinate]
			print(string(node.Height))
		}
		println()
	}
}

func (topMap *TopographyGraph) ResetExploredNodes() {
	for _, node := range topMap.NodeMap {
		node.IsExplored = false
		node.ParentNode = nil
	}
}

type Coordinates struct {
	X int
	Y int
}

func (c Coordinates) GetAdjacentSpace(dX int, dY int) Coordinates {
	return Coordinates{c.X + dX, c.Y + dY}
}

type MapNode struct {
	Height          int
	IsExplored      bool
	AccessibleNodes []*MapNode
	ParentNode      *MapNode
}

func (node *MapNode) IsAccessible(otherNode *MapNode) bool {
	return node.Height+1 >= otherNode.Height
}

func (node *MapNode) AppendAccessibleNode(aNode *MapNode) {
	node.AccessibleNodes = append(node.AccessibleNodes, aNode)
}
