package main

import "fmt"

type bicycler interface {
	spares() (int, *parts)
	defaultTireSize() int
	defaultChain() string
	localSpares() *parts
}

type bicycle struct {
	size     string
	chain    *parts
	tireSize int
}

func New(size string, chain *parts, tireSize int) *bicycle {
	return &bicycle{
		size:     size,
		chain:    chain,
		tireSize: tireSize,
	}
}
func (b *bicycle) spares() (int, *parts) {
	return b.tireSize, b.chain
}

func (b *bicycle) defaultTireSize() int {
	return 0
}

func (b *bicycle) defaultChain() string {
	return "11-speed"
}

func (b *bicycle) localSpares() *parts {
	return nil
}

func (rb *roadBike) defaultChain() string {
	return "13-speed"
}

type roadBike struct {
	bicycle
	tape_color string
}

type partItem [2]string
type parts []partItem

func (rb *roadBike) localSpares() *parts {

	return *rb.bicycle.spares() + partItem{}
}

//type parts []partItem
//type partItem struct {
//	name        string
//	description string
//	needsSpare  bool
//}

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
	//roadConfig := [][]string{
	//	{"chain", "11-speed"},
	//	{"tireSize", "23"},
	//	{"tapeColor", "red"},
	//}
	//bike = New("L", build(roadConfig), 10)

	//fmt.Printf("%v\n", roadBike.parts)
	var bike roadBike
	fmt.Println(bike.defaultChain())
}
