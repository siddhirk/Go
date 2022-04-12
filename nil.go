package main

import "fmt"

type Creature struct {
	Species string
}

func main() {
	var creature *Creature

	fmt.Printf("1) %+v\n", creature)
	changeCreature(creature)
	fmt.Printf("3) %+v\n", creature)
}

func changeCreature(creature *Creature) {
	creature.Species = "jellyfish"
	fmt.Printf("2) %+v\n", creature)
}
