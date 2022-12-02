package main

import (
	"log"
	"strings"
)

func CreateMatchList() []RPSMatch {
	var matches []RPSMatch
	beatMap := BeatMap()
	drawMap := DrawMap()

	onLineRead := func(line string) {
		parts := strings.Fields(line)
		opponent, err := CreateOption(beatMap, drawMap, parts[0])
		if err != nil {
			log.Printf("Could not create opponent option")
		}
		you, err := CreateOption(beatMap, drawMap, parts[1])
		if err != nil {
			log.Printf("Could not create your option")
		}

		match := RPSMatch{
			Opponent: opponent,
			You:      you,
		}

		matches = append(matches, match)
	}

	ReadFile("./resources/day_two_input.txt", onLineRead)

	return matches
}

func BeatMap() map[string]string {
	return map[string]string{
		"A": "Z",
		"B": "X",
		"C": "Y",
		"X": "C",
		"Y": "A",
		"Z": "B",
	}
}

func DrawMap() map[string]string {
	return map[string]string{
		"A": "X",
		"B": "Y",
		"C": "Z",
		"X": "A",
		"Y": "B",
		"Z": "C",
	}
}

func DayTwoA() int64 {
	var matches []RPSMatch
	beatMap := BeatMap()
	drawMap := DrawMap()

	onLineRead := func(line string) {
		parts := strings.Fields(line)
		opponent, err := CreateOption(beatMap, drawMap, parts[0])
		if err != nil {
			log.Printf("Could not create opponent option")
		}
		you, err := CreateOption(beatMap, drawMap, parts[1])
		if err != nil {
			log.Printf("Could not create your option")
		}

		match := RPSMatch{
			Opponent: opponent,
			You:      you,
		}

		matches = append(matches, match)
	}

	ReadFile("./resources/day_two_input.txt", onLineRead)

	var totalScore int64

	for _, match := range matches {
		totalScore = totalScore + match.CalculateScore()
	}

	return totalScore
}

func DayTwoB() int64 {
	var matches []RPSMatch
	beatMap := BeatMap()
	drawMap := DrawMap()

	onLineRead := func(line string) {
		parts := strings.Fields(line)
		opponent, err := CreateOption(beatMap, drawMap, parts[0])
		if err != nil {
			log.Printf("Could not create opponent option")
		}
		you, err := CreateCounterOption(beatMap, drawMap, opponent, parts[1])
		if err != nil {
			log.Printf("Could not create your option")
		}

		match := RPSMatch{
			Opponent: opponent,
			You:      you,
		}

		matches = append(matches, match)
	}

	ReadFile("./resources/day_two_input.txt", onLineRead)

	var totalScore int64

	for _, match := range matches {
		totalScore = totalScore + match.CalculateScore()
	}

	return totalScore
}

func DayTwo() {
	log.Print("Day 2:")
	dayTwoResult := DayTwoA()
	log.Printf("Answer A: %v", dayTwoResult)
	dayTwoBResult := DayTwoB()
	log.Printf("Answer B: %v", dayTwoBResult)
}
