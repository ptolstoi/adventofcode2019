package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main() {
	fmt.Printf(`
Day 2 --- Advent of Code 2019
-----------------------------
`)

	file := "./day2/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	rawData := strings.Split(string(data), ",")

	memory := make([]int, len(rawData))

	for i, k := range rawData {
		value, err := strconv.Atoi(k)
		if err != nil {
			log.Fatalf("couldn't parse value: %s", k)
		}
		memory[i] = value
	}

	childMemory := make([]int, len(memory))
	memory[1] = 12
	memory[2] = 2
	copy(childMemory, memory)
	fmt.Printf("part I  | value at position 0: %d\n", compute(childMemory))

mainLoop:
	for noun := 0; noun <= 99; noun++ {
		memory[1] = noun
		for verb := 0; verb <= 99; verb++ {
			memory[2] = verb

			copy(childMemory, memory)

			if compute(childMemory) == 19690720 {
				break mainLoop
			}
		}
	}

	fmt.Printf("part II | noun * 100 + verb: %d\n", memory[1]*100+memory[2])
}

func compute(memory []int) int {
	i := 0
mainLoop:
	for i < len(memory) {
		switch memory[i] {
		case 1:
			memory[memory[i+3]] = memory[memory[i+1]] + memory[memory[i+2]]
			i += 4
		case 2:
			memory[memory[i+3]] = memory[memory[i+1]] * memory[memory[i+2]]
			i += 4
		case 99:
			break mainLoop
		default:
			log.Fatalf("invalid opcode %v at %v", memory[i], i)
		}
	}
	if i >= len(memory) {
		log.Fatalf("index %d out of bounds!", i)
	}

	return memory[0]
}
