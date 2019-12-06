package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type intcoder struct {
	memory []int
	p      int
}

func Main() {
	fmt.Printf(`
Day 5 --- Advent of Code 2019
-----------------------------
`)

	file := "./day5/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	rawData := strings.Split(string(data), ",")

	programm := make([]int, len(rawData))

	for i, k := range rawData {
		value, err := strconv.Atoi(k)
		if err != nil {
			log.Fatalf("couldn't parse value: %s", k)
		}
		programm[i] = value
	}

	partI := newIntcoder(programm)
	partII := newIntcoder(programm)

	fmt.Printf("part I  | diagnostic value: %d\n", partI.compute(1))
	fmt.Printf("part II | diagnostic value: %d\n", partII.compute(5))
}

func newIntcoder(programm []int) intcoder {
	newCoder := intcoder{
		memory: make([]int, len(programm)),
	}
	copy(newCoder.memory, programm)

	return newCoder
}

func (_m *intcoder) compute(input int) int {
	lastOutput := _m.memory[0]
	_m.p = 0
mainLoop:
	for _m.p < len(_m.memory) {
		i := _m.p
		optcode := _m.memory[i]

		switch optcode % 100 {
		case 1: // add operator
			_m.memory[_m.memory[i+3]] = _m.extract(1) + _m.extract(2)
			_m.p += 4
		case 2: // multiply operator
			_m.memory[_m.memory[i+3]] = _m.extract(1) * _m.extract(2)
			_m.p += 4
		case 3: // input operator
			_m.memory[_m.memory[i+1]] = input
			_m.p += 2
		case 4: // output operator
			lastOutput = _m.extract(1)
			log.Printf("output: %v", lastOutput)
			_m.p += 2
		case 5: // jump-if-true
			if _m.extract(1) != 0 {
				_m.p = _m.extract(2)
			} else {
				_m.p += 3
			}
		case 6: // jump-if-false
			if _m.extract(1) == 0 {
				_m.p = _m.extract(2)
			} else {
				_m.p += 3
			}
		case 7: // is-less-than
			if _m.extract(1) < _m.extract(2) {
				_m.memory[_m.memory[i+3]] = 1
			} else {
				_m.memory[_m.memory[i+3]] = 0
			}
			_m.p += 4
		case 8: // equals
			if _m.extract(1) == _m.extract(2) {
				_m.memory[_m.memory[i+3]] = 1
			} else {
				_m.memory[_m.memory[i+3]] = 0
			}
			_m.p += 4
		case 99:
			break mainLoop
		default:
			log.Fatalf("invalid opcode %v at %v", _m.memory[i], i)
		}
	}
	if _m.p >= len(_m.memory) {
		log.Fatalf("index %d out of bounds!", _m.p)
	}

	return lastOutput
}

func (_m *intcoder) extract(index int) int {
	optcode := _m.memory[_m.p]
	mode := (optcode / p10(index+1)) % 10
	switch mode {
	case 0:
		return _m.memory[_m.memory[_m.p+index]]
	case 1:
		return _m.memory[_m.p+index]
	}

	log.Fatalf("invalid parameter mode %v for index %v: %v", mode, index, optcode)
	return -1
}

func p10(e int) int {
	if e == 0 {
		return 1
	}

	return 10 * p10(e-1)
}
