package day10

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"sort"
	"strings"
)

type vec struct {
	x, y int
}
type pos struct {
	x, y  int
	count int
	deg   float64

	reachable []pos
}

func Main() {
	fmt.Printf(`
Day 10 --- Advent of Code 2019
-----------------------------
`)

	file := "./day10/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	lines := strings.Split(string(data), "\n")
	grid := []pos{}

	for y, line := range lines {
		for x, sym := range line {
			if sym == '#' {
				grid = append(grid, pos{
					x:         x,
					y:         y,
					reachable: []pos{},
				})
			} else if sym == '.' {
			} else {
				log.Fatalf("invalid input at y=%v,x=%v: %v", y, x, sym)
			}
		}
	}

	maxIdx := -1

	for i, A := range grid {
		cnt := map[vec]struct{}{}
		for j, B := range grid {
			if i != j {
				xDiff := B.x - A.x
				yDiff := B.y - A.y

				diffGCD := gcd(xDiff, yDiff)

				key := vec{
					x: xDiff / diffGCD,
					y: yDiff / diffGCD,
				}

				if _, ok := cnt[key]; !ok {
					cnt[key] = struct{}{}
					grid[i].reachable = append(grid[i].reachable, B)
				}
			}
		}

		grid[i].count = len(cnt)
		if maxIdx < 0 || grid[maxIdx].count < grid[i].count {
			maxIdx = i
		}
	}

	station := grid[maxIdx]
	grid = station.reachable

	fmt.Printf("part I  | new station: x=%v y=%v cnt=%v\n", station.x, station.y, station.count)

	vStart := vec{x: 0, y: -1}
	for i := range grid {
		x := grid[i].x - station.x
		y := grid[i].y - station.y

		grid[i].deg =
			math.Atan2(float64(y), float64(x)) -
				math.Atan2(float64(vStart.y), float64(vStart.x))

		if grid[i].deg < 0 {
			grid[i].deg += 2 * math.Pi
		}
	}

	sort.SliceStable(grid, func(i, j int) bool {
		return grid[i].deg < grid[j].deg
	})

	last := grid[199]

	fmt.Printf("part II | new station: %v\n", last.x*100+last.y)
}

func gcd(a, b int) int {
	if a == 0 {
		return abs(b)
	}
	if b == 0 {
		return abs(a)
	}
	var h int
	for {
		h = a % b
		a = b
		b = h
		if b == 0 {
			break
		}
	}

	return abs(a)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
