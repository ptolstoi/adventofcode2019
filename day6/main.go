package day6

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type system map[string]string

func Main() {
	fmt.Printf(`
Day 6 --- Advent of Code 2019
-----------------------------
`)

	file := "./day6/input.data"

	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatalf("error when opening file %s: %v", file, err)
	}

	connections := strings.Split(string(data), "\n")

	systemConnections := system{}
	objectsMap := map[string]bool{}

	for _, connection := range connections {
		parts := strings.Split(connection, ")")
		if len(parts) != 2 {
			log.Fatalf("wrong connection: %v", connection)
		}

		systemConnections[parts[1]] = parts[0]

		objectsMap[parts[0]] = true
		objectsMap[parts[1]] = true
	}

	objects := make([]string, len(objectsMap))
	i := 0
	for object := range objectsMap {
		objects[i] = object
		i++
	}
	systemConnections.verify(objects)
	checksum := systemConnections.calcChecksum(objects)

	fmt.Printf("part I  | checksum: %d\n", checksum)

	youPath := systemConnections.getPathToCOM("YOU")
	santaPath := systemConnections.getPathToCOM("SAN")

	sameNodes := 0
	for youPath[sameNodes] == santaPath[sameNodes] {
		sameNodes++
	}
	fmt.Printf("part II | orbit transfers %d\n", (len(youPath)-sameNodes)+(len(santaPath)-sameNodes))
}

func (connections *system) verify(objects []string) {
	hasCOM := false

	for _, object := range objects {
		if object == "COM" {
			hasCOM = true
		} else if (*connections)[object] == "" {
			log.Fatalf("object %v has no parent! %v %v", object, connections, objects)
		}
	}

	if !hasCOM {
		log.Fatalf("COM missing!")
	}
}

func (connections *system) calcChecksum(objects []string) int {
	count := 0

	for _, object := range objects {
		count += len(connections.getPathToCOM(object))
	}

	return count
}

func (connections *system) getPathToCOM(object string) []string {
	if object == "COM" {
		return nil
	}

	parent, ok := (*connections)[object]
	if !ok {
		log.Fatalf("no parent found for %v: %v", object, connections)
	}

	return append(connections.getPathToCOM(parent), parent)
}
