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

func findClosestIntersectionDistance(pathA []string, pathB []string) (closestDistance int) {

	distance := findIntersectionDistance(pathA, pathB)

	sort.Sort(sort.IntSlice(distance))

	return distance[0]
}

func findIntersectionDistance(pathA []string, pathB []string) (distance []int) {

	intersections := findIntersections(pathA, pathB)

	for _, v := range intersections {
		distance = append(distance, int(math.Abs(float64(v.x)))+int(math.Abs(float64(v.y))))
	}

	return
}

func findIntersections(pathA []string, pathB []string) (intersections []vector) {

	g := make(map[vector]int)

	walkGrid(g, pathA, 1)
	walkGrid(g, pathB, 2)

	for k, v := range g {
		if v > 2 {
			intersections = append(intersections, k)
		}
	}

	return
}

func walkGrid(grid map[vector]int, path []string, pathID int) {

	var movements []movement

	for _, v := range path {
		d, _ := strconv.Atoi(v[1:len(v)])
		movements = append(movements, movement{v[0:1], d})
	}

	var offset vector = vector{0, 0}

	for _, v := range movements {
		offset = makeMove(grid, offset, v, pathID)
	}

}

func makeMove(grid map[vector]int, offset vector, movement movement, pathID int) vector {

	switch movement.direction {
	case "U":
		for i := 0; i < movement.distance; i++ {
			offset.y++
			if grid[offset] != pathID {
				grid[offset] = grid[offset] + pathID
			}
		}
	case "D":
		for i := 0; i < movement.distance; i++ {
			offset.y--
			if grid[offset] != pathID {
				grid[offset] = grid[offset] + pathID
			}
		}
	case "R":
		for i := 0; i < movement.distance; i++ {
			offset.x++
			if grid[offset] != pathID {
				grid[offset] = grid[offset] + pathID
			}
		}
	case "L":
		for i := 0; i < movement.distance; i++ {
			offset.x--
			if grid[offset] != pathID {
				grid[offset] = grid[offset] + pathID
			}
		}
	}

	return offset
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	var path [][]string
	for scanner.Scan() {
		movements := strings.Split(scanner.Text(), ",")

		path = append(path, movements)
	}

	distance := findClosestIntersectionDistance(path[0], path[1])

	fmt.Printf("Distance to closest intersection: %d", distance)

}
