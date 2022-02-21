package main

import "fmt"

type bicycle struct {
	size  string
	parts *parts
}

func NewBicycle(size string, config [][]string) *bicycle {
	return &bicycle{
		size:  size,
		parts: build(config),
	}
}

type parts []partItem
type partItem struct {
	name        string
	description string
	needsSpare  bool
}

type PartsMaker interface {
	build(config map[string]string, parts *parts) parts
}

func build(config [][]string) *parts {
	var parts parts
	for i := range config {
		parts = append(parts, createPart(config[i]))
	}
	return &parts
}

func createPart(partConfig []string) partItem {
	needsSpare := true
	if len(partConfig) > 2 && partConfig[2] == "false" {
		needsSpare = false
	}

	return partItem{
		name:        partConfig[0],
		description: partConfig[1],
		needsSpare:  needsSpare,
	}
}

func main() {
	roadConfig := [][]string{
		{"chain", "11-speed"},
		{"tireSize", "23"},
		{"tapeColor", "red"},
	}
	roadBike := NewBicycle("L", roadConfig)

	fmt.Printf("%v\n", roadBike.parts)
}
