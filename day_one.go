package main

import (
	"log"
	"sort"
	"strconv"
)

func DayOneA() Elf {
	elves := make([]Elf, 0)
	var elf = CreateElf()

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

	ReadFile(onLineRead)

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].Totalcalories() > elves[j].Totalcalories()
	})

	return elves[0]
}

func DayOneB() []Elf {
	elves := make([]Elf, 0)
	var elf = CreateElf()

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

	ReadFile(onLineRead)

	sort.SliceStable(elves, func(i, j int) bool {
		return elves[i].Totalcalories() > elves[j].Totalcalories()
	})

	return elves[0:3]
}
