package day8

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

const (
	black       int = 0
	white       int = 1
	transparent int = 2
)

func Main() {
	fmt.Printf(`
Day 8 --- Advent of Code 2019
-----------------------------
`)

	file := "./day8/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	width := 25
	height := 6

	if len(data)%(width*height) != 0 {
		log.Fatalf("data / %v / %v is not zero!", width, height)
	}

	layerCount := len(data) / width / height

	i := 0

	imageData := make([][][]int, layerCount)
	for l := range imageData {
		imageData[l] = make([][]int, height)
		for y := range imageData[l] {
			imageData[l][y] = make([]int, width)

			for x := range imageData[l][y] {
				v, err := strconv.Atoi(string(data[i]))
				if err != nil {
					log.Fatalf("couldn't convert %v: %v", string(data[i]), err)
				}
				imageData[l][y][x] = v
				i++
			}
		}
	}

	zeroCount := width * height
	layerWithFewestZeros := 0
	for i, layer := range imageData {
		count := count(layer, 0)
		if count < zeroCount {
			zeroCount = count
			layerWithFewestZeros = i
		}
	}

	checksum := count(imageData[layerWithFewestZeros], 1) * count(imageData[layerWithFewestZeros], 2)

	fmt.Printf("part I  | checksum: %d\n", checksum)
	fmt.Printf("part II | \n")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			finalColor := condense(imageData, x, y)
			switch finalColor {
			case transparent:
				fmt.Print(" ")
			case black:
				fmt.Print("⬛️")
			case white:
				fmt.Print("⬜️")
			default:

			}
		}
		fmt.Print("\n")
	}
}

func count(layer [][]int, val int) (result int) {
	for _, y := range layer {
		for _, x := range y {
			if x == val {
				result++
			}
		}
	}

	return
}

func condense(imageData [][][]int, x, y int) int {
	for l := range imageData {
		if imageData[l][y][x] == transparent {
			continue
		}

		return imageData[l][y][x]
	}

	return transparent
}
