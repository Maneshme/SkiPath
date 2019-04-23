package main

import (
	"fmt"
)

func main() {
	elevationMap := [][]int{
		{40, 88, 71, 32, 10},
		{21, 57, 94, 36, 8},
		{62, 39, 25, 58, 67},
		{48, 5, 10, 66, 39},
		{54, 2, 0, 31, 41},
	}

	bestPath := findBestSkiPath(elevationMap)
	fmt.Printf("The best ski path is: %+v\n", bestPath)
	fmt.Printf("Depth: %d", bestPath[0] - bestPath[len(bestPath) - 1])
}

// findBestSkiPath is to find the best ski path from the elevation map
func findBestSkiPath(elevationMap [][]int) []int {
	allPaths := make([][]int, 0)
	for hiIdx := 0; hiIdx < len(elevationMap); hiIdx++ {
		for wiIdx := 0; wiIdx < len(elevationMap[hiIdx]); wiIdx++ {
			allPaths = append(allPaths, findAllPaths(elevationMap[hiIdx][wiIdx], hiIdx, wiIdx, len(elevationMap), len(elevationMap[hiIdx]), elevationMap)...)
		}
	}
	path := make([]int, 0)
	if len(allPaths) != 0 {
		allPaths = reduceByLength(allPaths)
		fmt.Printf("Best Paths Found: %+v\n",allPaths)
		path = findSteepestPath(allPaths)
	}
	return path
}

// findSteepestPath is to find the steepest path in a list of paths
func findSteepestPath(allPaths [][]int) []int {
	maxDepth := 0
	steepest := make([]int, 0)
	for _, path := range allPaths {
		pathDepth :=  path[0] - path[len(path)-1]
		if pathDepth > maxDepth {
			maxDepth = pathDepth
			steepest = path
		}
	}
	return steepest;
}

// reduceByLength is to get all the paths with the maximum length
func reduceByLength(allPaths [][]int) [][]int {
	reducedPaths := make([][]int, 0)
	maxLength := getMaximumLength(allPaths)
	for _, path := range allPaths {
		if len(path) == maxLength {
			reducedPaths = append(reducedPaths, path)
		}
	}
	return reducedPaths
}

// getMaximumLength is to get the length of the longest path in a list of paths
func getMaximumLength(allPaths [][]int) int{
	maxLength := 0
	for idx := 0; idx < len(allPaths); idx++ {
		if len(allPaths[idx]) > maxLength {
			maxLength = len(allPaths[idx])
		}
	}
	return maxLength
}

// findAllPaths is find all paths available in the elevation map
func findAllPaths(start, hiIdx, wiIdx, height, width int, elevationMap [][]int) [][]int {
	allPaths := make([][]int, 0)
	if hiIdx != 0 &&
		elevationMap[hiIdx - 1][wiIdx] >= 0 &&
		elevationMap[hiIdx - 1][wiIdx] < start {
			northArr := []int{elevationMap[hiIdx - 1][wiIdx]}
			allPaths = append(allPaths, northArr)
			northPaths := findAllPaths(elevationMap[hiIdx - 1][wiIdx], hiIdx - 1, wiIdx, height, width, elevationMap)
			allPaths = append(allPaths, northPaths...)
	}

	if hiIdx + 1 < height &&
		elevationMap[hiIdx + 1][wiIdx] >= 0 &&
		elevationMap[hiIdx + 1][wiIdx] < start {
			southArr := []int{elevationMap[hiIdx + 1][wiIdx]}
			allPaths = append(allPaths, southArr)
			southPaths := findAllPaths(elevationMap[hiIdx + 1][wiIdx], hiIdx + 1, wiIdx, height, width, elevationMap)
			allPaths = append(allPaths, southPaths...)
	}

	if wiIdx != 0 &&
		elevationMap[hiIdx][wiIdx - 1] >= 0 &&
		elevationMap[hiIdx][wiIdx - 1] < start {
			westArr := []int{elevationMap[hiIdx][wiIdx - 1]}
			allPaths = append(allPaths, westArr)
			westPaths := findAllPaths(elevationMap[hiIdx][wiIdx - 1], hiIdx, wiIdx - 1, height, width, elevationMap)
			allPaths = append(allPaths, westPaths...)
	}

	if wiIdx + 1 < width &&
		elevationMap[hiIdx][wiIdx + 1] >= 0 &&
		elevationMap[hiIdx][wiIdx + 1] < start {
			eastArr := []int{elevationMap[hiIdx][wiIdx + 1]}
			allPaths = append(allPaths, eastArr)
			eastPaths := findAllPaths(elevationMap[hiIdx][wiIdx + 1], hiIdx, wiIdx + 1, height, width, elevationMap)
			allPaths = append(allPaths, eastPaths...)
	}
	allPaths = addStartToAllPaths(allPaths, start)
	return allPaths
}

func addStartToAllPaths(allPaths [][]int, start int) [][]int {
	for hiIdx := 0; hiIdx < len(allPaths); hiIdx++ {
		row := make([]int, 0)
		row = append(row, start)
		row = append(row, allPaths[hiIdx]...)
		allPaths[hiIdx] = row
	}
	return allPaths
}