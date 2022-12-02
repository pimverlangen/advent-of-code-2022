package main

import (
	"errors"
	"fmt"
)

type RPSMatch struct {
	Opponent RPSOption
	You      RPSOption
}

func (self *RPSMatch) CalculateScore() int64 {
	score := self.You.Value

	if self.You.Beats(self.Opponent) == true {
		score = score + 6
	} else if self.You.Draws(self.Opponent) {
		score = score + 3
	}

	return score
}

type RPSOption struct {
	Type  string
	Value int64
	Label string
	Beats func(against RPSOption) bool
	Draws func(against RPSOption) bool
}

func CreateOption(beatMap map[string]string, drawMap map[string]string, inputChar string) (RPSOption, error) {
	switch inputChar {
	case "A":
		return RPSOption{inputChar, 1, "Rock",
				func(against RPSOption) bool { return against.Type == beatMap[inputChar] },
				func(against RPSOption) bool { return against.Type == inputChar || against.Type == drawMap[inputChar] }},
			nil
	case "X":
		return RPSOption{inputChar, 1, "Rock Counter",
				func(against RPSOption) bool { return against.Type == beatMap[inputChar] },
				func(against RPSOption) bool { return against.Type == inputChar || against.Type == drawMap[inputChar] }},
			nil
	case "B":
		return RPSOption{inputChar, 2, "Paper",
				func(against RPSOption) bool { return against.Type == beatMap[inputChar] },
				func(against RPSOption) bool { return against.Type == inputChar || against.Type == drawMap[inputChar] }},
			nil
	case "Y":
		return RPSOption{inputChar, 2, "Paper Counter",
				func(against RPSOption) bool { return against.Type == beatMap[inputChar] },
				func(against RPSOption) bool { return against.Type == inputChar || against.Type == drawMap[inputChar] }},
			nil
	case "C":
		return RPSOption{inputChar, 3, "Scissor",
				func(against RPSOption) bool { return against.Type == beatMap[inputChar] },
				func(against RPSOption) bool { return against.Type == inputChar || against.Type == drawMap[inputChar] }},
			nil
	case "Z":
		return RPSOption{inputChar, 3, "Scissor Counter",
				func(against RPSOption) bool { return against.Type == beatMap[inputChar] },
				func(against RPSOption) bool { return against.Type == inputChar || against.Type == drawMap[inputChar] }},
			nil
	}

	return RPSOption{}, errors.New(fmt.Sprintf("Invalid inputChar [%v]", inputChar))
}

func CreateCounterOption(beatMap map[string]string, drawMap map[string]string, opposingOption RPSOption, expectedOutcome string) (RPSOption, error) {
	if expectedOutcome == "X" {
		// lose
		for winner, loser := range beatMap {
			if winner == opposingOption.Type {
				return CreateOption(beatMap, drawMap, loser)
			}
		}
	} else if expectedOutcome == "Y" {
		// draw
		return CreateOption(beatMap, drawMap, opposingOption.Type)
	} else {
		// win
		for winner, loser := range beatMap {
			if loser == opposingOption.Type {
				return CreateOption(beatMap, drawMap, winner)
			}
		}
	}

	return RPSOption{}, errors.New(fmt.Sprintf("Could not create counter option [opposingOption: %v]", opposingOption))
}
