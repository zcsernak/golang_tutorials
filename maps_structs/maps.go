package main

import (
	"fmt"
	"reflect"
)

func main() {
	// Maps
	fmt.Println("Maps")
	// Liter syntax          key    value
	statePopulations := map[string]int{
		"California":  39250017,
		"Texas":       27862596,
		"Florida":     20612439,
		"New York":    19745289,
		"Pensylvania": 12802503,
		"Illinois":    12801539,
		"Ohio":        11614373,
	}
	fmt.Println(statePopulations)

	m := map[[3]int]string{} // invalid map key type []int <- slices can't be keys, but arrays can be
	fmt.Println(m)

	statePopulations2 := make(map[string]int)
	statePopulations2 = map[string]int{
		"Budapest": 1756000,
	}
	fmt.Println(statePopulations2)

	fmt.Println(statePopulations["Ohio"])
	statePopulations["Georgia"] = 10310371
	fmt.Println(statePopulations) // return order of a map is not guaranteed

	delete(statePopulations, "Georgia")
	fmt.Println(statePopulations)

	pop, ok := statePopulations["Georgia"]
	fmt.Println(pop, ok) // 'ok' is false if the key is not in the map

	fmt.Println(len(statePopulations))

	// maps are reference types (they are passed by reference)
	sp := statePopulations
	delete(sp, "Ohio")
	fmt.Println(sp)
	fmt.Println(statePopulations)

	StructsFunc()
	CompositionFunc()
	TagFunc()
}

// exported from the package, becaue of capital first letter,
// but field names aren't exported because of lowercase first letters
type Doctor struct {
	number     int
	actorName  string
	companions []string
	episodes   []string
}

func StructsFunc() {
	fmt.Println("\nStucts")
	aDoctor := Doctor{
		number:    3, // field names are optional, but if they are left out they can be a maintenance problem when adding new fields
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Jane Smith",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[1])

	// anonymous struct
	bDoctor := struct{ name string }{name: "J.D."}
	fmt.Println(bDoctor)

	// structs are value types (they are copied when passed)
	anotherDoctor := bDoctor
	anotherDoctor.name = "Turk"
	fmt.Println(anotherDoctor)
	fmt.Println(bDoctor)
	pointerDoc := &anotherDoctor
	pointerDoc.name = "Dr. Cox"
	fmt.Println(anotherDoctor)
}

type Animal struct {
	Name   string
	Origin string
}

type Creature struct {
	Name        string
	AttackValue int32
}

type Bird struct {
	Animal // embed the Animal struct in the Bird struct
	Creature
	SpeedKPM float32
	CanFly   bool
}

func CompositionFunc() {
	fmt.Println("\nComposition (embedding structs)")
	b := Bird{}
	b.Animal.Name = "Emu"
	b.Creature.Name = "Battle Bird"
	b.Origin = "Australia"
	b.AttackValue = 9000
	b.SpeedKPM = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Animal.Name)
	fmt.Println(b.Creature.Name)

	c := Bird{
		Animal:   Animal{Name: "Penguin", Origin: "North Pole"},
		SpeedKPM: 36,
		CanFly:   false,
	}
	fmt.Println(c)

}

type Monster struct {
	Name   string `required max:"100"`
	Origin string
}

func TagFunc() {
	fmt.Println("\nTag")
	t := reflect.TypeOf(Monster{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // validation framework/library is needed to parse this
}
