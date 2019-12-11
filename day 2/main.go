package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func parseIntcode(intCode string, noun string, verb string) []string {
	memory := strings.Split(intCode, ",")

	memory[1] = noun
	memory[2] = verb

	instructionPointer := 0

	for (instructionPointer < len(memory)) && (parseInstruction(memory, instructionPointer) >= 0) {
		instructionPointer = instructionPointer + 4
	}

	return memory
}

func parseInstruction(intCode []string, pos int) int {
	opcode := intCode[pos]
	if opcode == "1" {
		a, b, c := getInstructionParameters(intCode, pos)
		intCode[c] = strconv.Itoa(a + b)
		return 0
	} else if opcode == "2" {
		a, b, c := getInstructionParameters(intCode, pos)
		intCode[c] = strconv.Itoa(a * b)
		return 0
	} else if opcode == "99" {
		return -1
	}
	return -1
}

func getInstructionParameters(intCode []string, pos int) (a int, b int, c int) {
	a, _ = strconv.Atoi(intCode[pos+1])
	a, _ = strconv.Atoi(intCode[a])
	b, _ = strconv.Atoi(intCode[pos+2])
	b, _ = strconv.Atoi(intCode[b])
	c, _ = strconv.Atoi(intCode[pos+3])
	return
}

func main() {
	body, _ := ioutil.ReadFile("input.txt")
	buffer := bytes.NewBuffer(body)
	for _, m := range strings.Split(buffer.String(), "\n") {
		fmt.Printf("%s\n", parseIntcode(m, "12", "2"))
		for noun := 0; noun <= 99; noun++ {
			for verb := 0; verb <= 99; verb++ {
				if parseIntcode(m, strconv.Itoa(noun), strconv.Itoa(verb))[0] == "19690720" {
					fmt.Printf("Pair that produces output: 19690720 found: noun: %d, verb: %d\n", noun, verb)
					fmt.Printf("Solution is %d", 100*noun+verb)
				}
			}
		}
	}

}
