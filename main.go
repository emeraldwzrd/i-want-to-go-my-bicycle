package main

type bicycle struct {
	size  int
	parts parts
}

func (b *bicycle) initialize(size int, parts parts) {
	b.size = size
	b.parts = parts
}

func (b *bicycle) spares(size int, parts parts) {
	b.size = size
	b.parts = parts
}

type parts struct {
	spares int
	size   int
	parts  int
}

func (p *parts) initialize(parts int) {
	p.parts = parts
}

//func (p *parts) spares(parts int) {
//	p.parts = parts
//}

type partsfactory interface {
	//name        string
	//description string
	//needsSpare  string
}

func build(config []int) parts {
	return parts{
		spares: 0,
		size:   0,
		parts:  0,
	}
}

func createPart(partConfig []string) map[string]string {
	pf := make(map[string]string)
	pf["name"] = partConfig[0]
	pf["description"] = partConfig[1]
	pf["needsSpare"] = partConfig[2]
	return pf
}
