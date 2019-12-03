package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ptolstoi/adventofcode2019/day1"
)

func main() {
	day := flag.Int("day", 0, "day")
	flag.Parse()

	switch *day {
	case 1:
		day1.Main()
	default:
		fmt.Printf("Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}
}
