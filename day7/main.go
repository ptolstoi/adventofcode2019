package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type intcoder struct {
	programm []int
	memory   []int
	p        int
}

func Main() {
	fmt.Printf(`
Day 7 --- Advent of Code 2019
-----------------------------
`)

	file := "./day7/input.data"

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

	amplifiers := make([]*intcoder, 5)
	for i := 0; i < len(amplifiers); i++ {
		amplifiers[i] = newIntcoder(programm)
	}

	maxOutput := 0
	phase := []int{}

	allPhaseSettings := permutate([]int{0, 1, 2, 3, 4})
	for _, phaseSettings := range allPhaseSettings {
		output := runAmplification(amplifiers, phaseSettings)
		if output > maxOutput {
			maxOutput = output
			phase = phaseSettings
		}
	}

	fmt.Printf("part I  | diagnostic value: %d %v\n", maxOutput, phase)

	allPhaseSettingsLoop := permutate([]int{5, 6, 7, 8, 9})

	maxOutput = 0
	phase = []int{}
	for _, phaseSettings := range allPhaseSettingsLoop {
		output := runAmplificationLoop(amplifiers, phaseSettings)
		if output > maxOutput {
			maxOutput = output
			phase = phaseSettings
		}
	}

	fmt.Printf("part II | diagnostic value: %d %v\n", maxOutput, phase)
}

func newIntcoder(programm []int) *intcoder {
	newCoder := intcoder{
		memory:   make([]int, len(programm)),
		programm: make([]int, len(programm)),
	}
	copy(newCoder.programm, programm)

	return &newCoder
}

func (_m *intcoder) reset() {
	copy(_m.memory, _m.programm)
	_m.p = -1
}

func (_m *intcoder) isHalted() bool {
	return _m.memory[_m.p] == 99
}

func (_m *intcoder) isReset() bool {
	return _m.p < 0
}

func (_m *intcoder) compute(input, output chan int) {
	if _m.p < 0 {
		_m.p = 0
	}
mainLoop:
	for _m.p < len(_m.memory) {
		optcode := _m.memory[_m.p]

		switch optcode % 100 {
		case 1: // add operator
			_m.memory[_m.memory[_m.p+3]] = _m.extract(1) + _m.extract(2)
			_m.p += 4
		case 2: // multiply operator
			_m.memory[_m.memory[_m.p+3]] = _m.extract(1) * _m.extract(2)
			_m.p += 4
		case 3: // input operator
			select {
			case val := <-input:
				_m.memory[_m.memory[_m.p+1]] = val
				_m.p += 2
			default:
				return
			}
		case 4: // output operator
			lastOutput := _m.extract(1)
			_m.p += 2
			output <- lastOutput
			if _m.memory[_m.p] == 99 {
				return
			}
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
				_m.memory[_m.memory[_m.p+3]] = 1
			} else {
				_m.memory[_m.memory[_m.p+3]] = 0
			}
			_m.p += 4
		case 8: // equals
			if _m.extract(1) == _m.extract(2) {
				_m.memory[_m.memory[_m.p+3]] = 1
			} else {
				_m.memory[_m.memory[_m.p+3]] = 0
			}
			_m.p += 4
		case 99:
			break mainLoop
		default:
			log.Fatalf("invalid opcode %v at %v", _m.memory[_m.p], _m.p)
		}
	}
	if _m.p >= len(_m.memory) {
		log.Fatalf("index %d out of bounds!", _m.p)
	}
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

func runAmplification(amplifiers []*intcoder, phaseSettings []int) int {
	input := make(chan int, 2)
	output := make(chan int, 1)

	for i, amplifier := range amplifiers {
		amplifier.reset()
		input <- phaseSettings[i]
		if i == 0 {
			input <- 0
		} else {
			input <- <-output
		}
		amplifier.compute(input, output)
	}

	return <-output
}

func runAmplificationLoop(amplifiers []*intcoder, phaseSettings []int) int {
	for _, amplifier := range amplifiers {
		amplifier.reset()
	}
	firstTime := true

	input := make(chan int, 2)
	output := make(chan int, 1)

	for {

		for i, amplifier := range amplifiers {
			x := amplifier.p
			if x < 0 {
				x = 0
			}

			if amplifier.isReset() {
				input <- phaseSettings[i]
			}
			if firstTime {
				input <- 0
				firstTime = false
			} else {
				val := <-output
				input <- val
			}

			amplifier.compute(input, output)
		}

		if amplifiers[len(amplifiers)-1].isHalted() {
			return <-output
		}
	}
}

func permutate(candidates []int) [][]int {
	if len(candidates) == 1 {
		return [][]int{candidates}
	}

	output := [][]int{}
	for i, candidate := range candidates {
		other := []int{}
		other = append(other, candidates[0:i]...)
		other = append(other, candidates[i+1:]...)
		allOthers := permutate(other)

		for _, other := range allOthers {
			output = append(output, append([]int{candidate}, other...))
		}
	}

	return output
}
