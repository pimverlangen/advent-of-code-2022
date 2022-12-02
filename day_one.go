package main

import (
	"log"
	"sort"
	"strconv"
)

func CreateElfList() []Elf {
	var elves []Elf
	elf := CreateElf()

	onLineRead := func(line string) {
		if line == "" {
			elves = append(elves, elf)
			elf = CreateElf()
		} else {
			calories, err := strconv.ParseInt(line, 10, 0)
			if err != nil {
				log.Print("Skipping elf, invalid calories")
			} else {
				elf.items = append(elf.items, ElfItem{
					calories: calories,
				})
			}
		}
	}

	ReadFile("./resources/day_one_input.txt", onLineRead)

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].Totalcalories() > elves[j].Totalcalories()
	})

	return elves
}

func DayOneA() Elf {
	elves := CreateElfList()

	return elves[0]
}

func DayOneB() []Elf {
	elves := CreateElfList()

	return elves[0:3]
}

func DayOne() {
	log.Print("Day 1")
	dayOneAResult := DayOneA()
	log.Printf("Answer A: %v", dayOneAResult.Totalcalories())
	dayOneBResult := DayOneB()
	log.Printf("Answer B: %v", Totalcalories(dayOneBResult))
}
