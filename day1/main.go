package day1

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func Main() {
	fmt.Printf(`
Day 1 --- Advent of Code 2019
-----------------------------
`)

	file := "./day1/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	inputs := strings.Split(string(data), "\n")

	modules := make([]int, len(inputs))
	for i, input := range inputs {
		moduleMass, err := strconv.Atoi(input)
		if err != nil {
			log.Fatalf("module %d is not a valid number: %v", i, err)
		}
		modules[i] = moduleMass
	}

	totalFuelNeeded := 0
	fuelNeededForModulesOnly := 0

	for _, moduleMass := range modules {
		// part I
		fuelNeededForModulesOnly += moduleMass/3 - 2
		// part II
		totalFuelNeeded += calcFuelForMass(moduleMass)
	}

	fmt.Printf("part I  | total needed fuel: %d\n", fuelNeededForModulesOnly)
	fmt.Printf("part II | total needed fuel: %d\n", totalFuelNeeded)
}

func calcFuelForMass(mass int) (fuel int) {
	for {
		mass = mass/3 - 2 // int division -> automatic rounding down
		if mass > 0 {
			fuel += mass
		} else {
			return
		}
	}
}
