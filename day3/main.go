package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type vector struct {
	x, y   int
	length uint
}

func Main() {
	fmt.Printf(`
Day 3 --- Advent of Code 2019
-----------------------------
`)

	file := "./day3/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	rawData := strings.Split(string(data), "\n")
	wires := make([][]vector, len(rawData))
	for i, wire := range rawData {
		for _, dir := range strings.Split(wire, ",") {
			length, err := strconv.Atoi(dir[1:])
			if err != nil {
				log.Fatalf("error when parsing vector %s: %v", dir, err)
			}
			x, y := 0, 0
			switch rune(dir[0]) {
			case 'U':
				y = 1
			case 'D':
				y = -1
			case 'L':
				x = -1
			case 'R':
				x = 1
			default:
				log.Fatalf("unknown direction: %v", dir)
			}

			wires[i] = append(wires[i], vector{x, y, uint(length)})
		}
	}

	distance := 0
	distanceWireLength := 0
	grids := map[int]map[int]map[int]int{}

	for wireID, wire := range wires {
		wireLength := 0

		grids[wireID] = map[int]map[int]int{}

		x := 0
		y := 0

		for _, dir := range wire {
			var j uint
			for j = 0; j < dir.length; j++ {
				wireLength++
				x += dir.x
				y += dir.y
				if grids[wireID][x] == nil {
					grids[wireID][x] = make(map[int]int)
				}
				grids[wireID][x][y] = wireLength

				for k := 0; k < wireID; k++ {
					if grids[k][x][y] != 0 {
						dist := abs(x) + abs(y)
						if distance == 0 || distance > dist {
							distance = dist
						}

						dist = wireLength + grids[k][x][y]
						if distanceWireLength == 0 || distanceWireLength > dist {
							distanceWireLength = dist
						}
					}
				}
			}
		}
	}

	fmt.Printf("part I  | distance: %v\n", distance)
	fmt.Printf("part II | distance: %v\n", distanceWireLength)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
