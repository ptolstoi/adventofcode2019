package day9

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type intcoder struct {
	memory   []int64
	programm []int64
	p        int
	r        int
}

func Main() {
	fmt.Printf(`
Day 9 --- Advent of Code 2019
-----------------------------
`)

	file := "./day9/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	rawData := strings.Split(string(data), ",")

	programm := make([]int64, len(rawData))

	for i, k := range rawData {
		value, err := strconv.ParseInt(k, 10, 64)
		if err != nil {
			log.Fatalf("couldn't parse value: %s", k)
		}
		programm[i] = value
	}

	boost := newIntcoder(programm)
	boost.reset()

	fmt.Printf("part I  | diagnostic value: %d\n", boost.compute(1))
	boost.reset()
	fmt.Printf("part II | diagnostic value: %d\n", boost.compute(2))
}

func newIntcoder(programm []int64) intcoder {
	newCoder := intcoder{
		programm: make([]int64, len(programm)+2000),
		memory:   make([]int64, len(programm)+2000),
	}
	copy(newCoder.programm, programm)

	return newCoder
}

func (_m *intcoder) reset() {
	_m.p = 0
	_m.r = 0
	copy(_m.memory, _m.programm)
}

func (_m *intcoder) compute(input int64) int64 {
	lastOutput := _m.memory[0]
mainLoop:
	for _m.p < len(_m.memory) {
		i := _m.p
		optcode := _m.memory[i]

		switch optcode % 100 {
		case 1: // add operator
			target := _m.target(3, true)
			_m.memory[target] = _m.extract(1) + _m.extract(2)
			_m.p += 4
		case 2: // multiply operator
			target := _m.target(3, true)
			_m.memory[target] = _m.extract(1) * _m.extract(2)
			_m.p += 4
		case 3: // input operator
			target := _m.target(1, true)
			_m.memory[target] = input
			_m.p += 2
		case 4: // output operator
			lastOutput = _m.extract(1)
			log.Printf("output: %v", lastOutput)
			_m.p += 2
		case 5: // jump-if-true
			if _m.extract(1) != 0 {
				_m.p = int(_m.extract(2))
			} else {
				_m.p += 3
			}
		case 6: // jump-if-false
			if _m.extract(1) == 0 {
				_m.p = int(_m.extract(2))
			} else {
				_m.p += 3
			}
		case 7: // is-less-than
			target := _m.target(3, true)
			if _m.extract(1) < _m.extract(2) {
				_m.memory[target] = 1
			} else {
				_m.memory[target] = 0
			}
			_m.p += 4
		case 8: // equals
			target := _m.target(3, true)
			if _m.extract(1) == _m.extract(2) {
				_m.memory[target] = 1
			} else {
				_m.memory[target] = 0
			}
			_m.p += 4
		case 9: // set relative base
			_m.r += int(_m.extract(1))
			_m.p += 2
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

func (_m *intcoder) target(index int, write bool) int {
	optcode := _m.memory[_m.p]
	mode := (optcode / p10(index+1)) % 10
	if write && mode == 1 {
		mode = 0
	}
	switch mode {
	case 0:
		return int(_m.memory[_m.p+index])
	case 1:
		return _m.p + index
	case 2:
		return _m.r + int(_m.memory[_m.p+index])
	}

	log.Fatalf("invalid parameter mode %v for index %v: %v", mode, index, optcode)
	return -1
}

func (_m *intcoder) extract(index int) int64 {
	target := _m.target(index, false)

	return _m.memory[target]
}

func p10(e int) int64 {
	if e == 0 {
		return 1
	}

	return 10 * p10(e-1)
}
