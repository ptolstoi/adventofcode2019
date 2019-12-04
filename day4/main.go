package day4

import (
	"fmt"
)

type vector struct {
	x, y   int
	length uint
}

func Main() {
	fmt.Printf(`
Day 4 --- Advent of Code 2019
-----------------------------
`)

	min := 235741
	max := 706948

	countI := 0
	countII := 0

	d := make([]int, 6)
	for password := min; password <= max; password++ {
		for i := 0; i < 6; i++ {
			d[5-i] = password % (pow(10, (i + 1))) / pow(10, i)
		}

		if d[0] > d[1] || d[1] > d[2] || d[2] > d[3] || d[3] > d[4] || d[4] > d[5] {
			continue
		}

		if d[0] == d[1] || d[1] == d[2] || d[2] == d[3] || d[3] == d[4] || d[4] == d[5] {
		} else {
			continue
		}

		countI++

		if d[0] == d[1] && d[1] != d[2] ||
			d[0] != d[1] && d[1] == d[2] && d[2] != d[3] ||
			d[1] != d[2] && d[2] == d[3] && d[3] != d[4] ||
			d[2] != d[3] && d[3] == d[4] && d[4] != d[5] ||
			d[3] != d[4] && d[4] == d[5] {

		} else {
			continue
		}

		countII++
	}

	fmt.Printf("part I  | password count: %v\n", countI)
	fmt.Printf("part II | password count: %v\n", countII)
}

func pow(i int, pwr int) (result int) {
	if pwr == 0 {
		return 1
	}

	return i * pow(i, pwr-1)
}
