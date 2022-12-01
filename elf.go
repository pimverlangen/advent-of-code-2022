package main

type ElfItem struct {
	calories int64
}

type Elf struct {
	items []ElfItem
}

func (self *Elf) Totalcalories() int64 {
	var totalCalories int64 = 0
	for _, item := range self.items {
		totalCalories = totalCalories + item.calories
	}
	return totalCalories
}

func Totalcalories(elves []Elf) int64 {
	var totalCalories int64 = 0
	for _, elf := range elves {
		totalCalories = totalCalories + elf.Totalcalories()
	}
	return totalCalories
}

func CreateElf() Elf {
	return Elf{
		items: []ElfItem{},
	}
}
