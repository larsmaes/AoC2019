package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type vector struct {
	x int
	y int
}

type movement struct {
	direction string
	distance  int
}

type wire struct {
	crossing int
	step     int
}

func findLowestIntersectionStepCount(paths [][]string) (lowestStepcount int) {
	intersections := findIntersections(paths[0], paths[1])

	var stepcount []int

	for _, v := range intersections {
		stepcount = append(stepcount, v)
	}

	sort.Sort(sort.IntSlice(stepcount))

	return stepcount[0]
}

func findClosestIntersectionDistance(paths [][]string) (closestDistance int) {

	distance := findIntersectionDistance(paths[0], paths[1])

	sort.Sort(sort.IntSlice(distance))

	return distance[0]
}

func findIntersectionDistance(pathA []string, pathB []string) (distance []int) {

	intersections := findIntersections(pathA, pathB)

	for k, _ := range intersections {
		distance = append(distance, int(math.Abs(float64(k.x)))+int(math.Abs(float64(k.y))))
	}

	return
}

func findIntersections(pathA []string, pathB []string) (intersections map[vector]int) {

	//a := make(map[string]int)
	g := make(map[vector]map[string]int)

	walkGrid(g, pathA, 1)
	walkGrid(g, pathB, 2)

	intersections = make(map[vector]int)

	for k, v := range g {
		if v["wire"] > 2 {
			intersections[k] = v["step"]
		}
	}

	return
}

func walkGrid(grid map[vector]map[string]int, path []string, pathID int) {

	var movements []movement

	for _, v := range path {
		d, _ := strconv.Atoi(v[1:len(v)])
		movements = append(movements, movement{v[0:1], d})
	}

	var offset vector = vector{0, 0}
	stepoffset := 1

	for _, v := range movements {
		offset, stepoffset = makeMove(grid, offset, v, pathID, stepoffset)
	}

}

func makeMove(grid map[vector]map[string]int, offset vector, movement movement, pathID int, stepoffset int) (vector, int) {

	var i int
	for i = 0; i < movement.distance; i++ {
		switch movement.direction {
		case "U":
			offset.y++
			a := grid[offset]
			if a == nil {
				a = make(map[string]int)
				grid[offset] = a
			}
			addWire(grid[offset], pathID, i, stepoffset)
		case "D":
			offset.y--
			a := grid[offset]
			if a == nil {
				a = make(map[string]int)
				grid[offset] = a
			}
			addWire(grid[offset], pathID, i, stepoffset)
		case "R":
			offset.x++
			a := grid[offset]
			if a == nil {
				a = make(map[string]int)
				grid[offset] = a
			}
			addWire(grid[offset], pathID, i, stepoffset)
		case "L":
			offset.x--
			a := grid[offset]
			if a == nil {
				a = make(map[string]int)
				grid[offset] = a
			}
			addWire(grid[offset], pathID, i, stepoffset)
		}

	}

	stepoffset += i

	return offset, stepoffset
}

func addWire(gridpoint map[string]int, pathID int, step int, stepoffset int) {
	if gridpoint["wire"] != pathID {
		gridpoint["wire"] += pathID
		gridpoint["step"] += step + stepoffset
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var paths [][]string
	for scanner.Scan() {
		movements := strings.Split(scanner.Text(), ",")

		paths = append(paths, movements)
	}

	distance := findClosestIntersectionDistance(paths)
	stepcount := findLowestIntersectionStepCount(paths)

	fmt.Printf("Distance to closest intersection: %d\n", distance)
	fmt.Printf("Lowest stepcount: %d\n", stepcount)

}
