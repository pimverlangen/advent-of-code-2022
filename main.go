package main

import (
	"log"
)

func main() {
	dayOneAResult := DayOneA()
	log.Printf("Result day one A: elf: [%v], total calories: [%v]", dayOneAResult, dayOneAResult.Totalcalories())
	dayOneBResult := DayOneB()
	log.Printf("Result day one B: elves: [%v], total calories: [%v]", dayOneBResult, Totalcalories(dayOneBResult))
}
