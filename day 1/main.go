package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func calcNeededFuel(mass int) int {
	return mass/3 - 2
}

func calcTotalModuleFuel(mass int) int {
	neededFuel := calcNeededFuel(mass)
	totalNeededFuel := neededFuel
	for neededFuel > 0 {
		neededFuel = calcNeededFuel(neededFuel)
		if neededFuel > 0 {
			totalNeededFuel += neededFuel
		}
	}
	return totalNeededFuel
}

func main() {
	body, _ := ioutil.ReadFile("input.txt")
	buffer := bytes.NewBuffer(body)
	var totalFuel int
	for _, m := range strings.Split(buffer.String(), "\n") {
		s, _ := strconv.Atoi(m)
		totalFuel += calcTotalModuleFuel(s)
	}
	fmt.Printf("Total fuel needed: %d\n", totalFuel)
}
