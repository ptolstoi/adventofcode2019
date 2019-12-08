package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ptolstoi/adventofcode2019/day1"
	"github.com/ptolstoi/adventofcode2019/day2"
	"github.com/ptolstoi/adventofcode2019/day3"
	"github.com/ptolstoi/adventofcode2019/day4"
	"github.com/ptolstoi/adventofcode2019/day5"
	"github.com/ptolstoi/adventofcode2019/day6"
	"github.com/ptolstoi/adventofcode2019/day7"
	"github.com/ptolstoi/adventofcode2019/day8"
)

func main() {
	day := flag.Int("day", 0, "day")
	flag.Parse()

	switch *day {
	case 1:
		day1.Main()
	case 2:
		day2.Main()
	case 3:
		day3.Main()
	case 4:
		day4.Main()
	case 5:
		day5.Main()
	case 6:
		day6.Main()
	case 7:
		day7.Main()
	case 8:
		day8.Main()
	default:
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}
