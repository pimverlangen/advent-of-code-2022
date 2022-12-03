package main

import (
	"log"
	"unicode"
)

type Item struct {
	Character rune
	priority  int64
}

type Compartment struct {
	items map[rune]Item
}

type Rucksack struct {
	AllItems         []Item
	leftCompartment  Compartment
	rightCompartment Compartment
}

type RucksackGroup struct {
	Rucksacks []Rucksack
}

func PriorityMapLowercase() map[rune]int64 {
	return map[rune]int64{
		'a': 1,
		'b': 2,
		'c': 3,
		'd': 4,
		'e': 5,
		'f': 6,
		'g': 7,
		'h': 8,
		'i': 9,
		'j': 10,
		'k': 11,
		'l': 12,
		'm': 13,
		'n': 14,
		'o': 15,
		'p': 16,
		'q': 17,
		'r': 18,
		's': 19,
		't': 20,
		'u': 21,
		'v': 22,
		'w': 23,
		'x': 24,
		'y': 25,
		'z': 26,
	}
}

func PriorityMapUppercase(priorityMapLowercase map[rune]int64) map[rune]int64 {
	priorityMap := map[rune]int64{}
	for k, v := range priorityMapLowercase {
		priorityMap[unicode.ToUpper(k)] = v + 26
	}
	return priorityMap
}

func PriorityMap() map[rune]int64 {
	// creates lowercase map
	priorityMapLowercase := PriorityMapLowercase()

	// creates uppercase map from lowercase map
	priorityMapUppercase := PriorityMapUppercase(priorityMapLowercase)

	// adds uppercase map to lowercase map
	for k, v := range priorityMapUppercase {
		priorityMapLowercase[k] = v
	}

	// returns combined uppercase + lowercase map
	return priorityMapLowercase
}

func CreateItem(char rune, priorityMap map[rune]int64) Item {
	return Item{
		Character: char,
		priority:  priorityMap[char],
	}
}

func (self *Rucksack) FindCommonItemInSack() Item {
	var commonItem Item
	for k, item := range self.leftCompartment.items {
		if _, ok := self.rightCompartment.items[k]; ok {
			commonItem = item
			break
		}
	}

	return commonItem
}

func (self *Rucksack) Contains(item Item) bool {
	for _, v := range self.AllItems {
		if v.Character == item.Character {
			return true
		}
	}

	return false
}

func (self *Rucksack) FindCommonItemInOtherSacks(others []Rucksack) Item {
	var commonItem Item
	for _, item := range self.AllItems {
		findCount := 0
		for _, sack := range others {
			inSack := sack.Contains(item)
			if inSack {
				findCount = findCount + 1
			}
		}

		if findCount == len(others) {
			commonItem = item
		}
	}
	return commonItem
}

func (self *RucksackGroup) FindCommonItemInGroup() Item {
	var commonItem Item
	for _, rucksack := range self.Rucksacks {
		commonItem = rucksack.FindCommonItemInOtherSacks(self.Rucksacks)
	}

	return commonItem
}

func CreateRucksacks() []Rucksack {
	priorityMap := PriorityMap()

	var rucksacks []Rucksack

	onLineRead := func(line string) {
		leftCompartment := Compartment{items: make(map[rune]Item)}
		rightCompartment := Compartment{items: make(map[rune]Item)}

		leftCompartmentString := line[0 : len(line)/2]
		rightCompartmentString := line[len(line)/2:]

		var allItems []Item

		for _, ch := range leftCompartmentString {
			item := CreateItem(ch, priorityMap)
			allItems = append(allItems, item)
			leftCompartment.items[ch] = item
		}

		for _, ch := range rightCompartmentString {
			item := CreateItem(ch, priorityMap)
			allItems = append(allItems, item)
			rightCompartment.items[ch] = item
		}

		rucksack := Rucksack{
			AllItems:         allItems,
			leftCompartment:  leftCompartment,
			rightCompartment: rightCompartment,
		}

		rucksacks = append(rucksacks, rucksack)
	}

	ReadFile("./resources/day_three_input.txt", onLineRead)

	return rucksacks
}

func DayThreeA() int64 {
	rucksacks := CreateRucksacks()

	var sumPriorities int64

	for _, sack := range rucksacks {
		sumPriorities = sumPriorities + sack.FindCommonItemInSack().priority
	}

	return sumPriorities
}

func DayThreeB() int64 {
	rucksacks := CreateRucksacks()

	var groups []RucksackGroup

	groupSize := 3
	var j int
	for i := 0; i < len(rucksacks); i += groupSize {
		j += groupSize
		if j > len(rucksacks) {
			j = len(rucksacks)
		}

		groups = append(groups, RucksackGroup{Rucksacks: rucksacks[i:j]})
	}

	var sumPriorities int64

	for _, group := range groups {
		sumPriorities = sumPriorities + group.FindCommonItemInGroup().priority
	}

	return sumPriorities
}

func DayThree() {
	log.Print("Day 3")
	dayThreeAResult := DayThreeA()
	log.Printf("Result A: %v", dayThreeAResult)
	dayThreeBResult := DayThreeB()
	log.Printf("Result B: %v", dayThreeBResult)
}
