package main

import (
	"log"
	"strconv"
	"strings"
)

type Section struct {
	start int64
	end   int64
}

func (self *Section) HasCompleteOverlapWith(other Section) bool {
	return self.start <= other.start && self.end >= other.end
}

func (self *Section) HasPartialOverlapWith(other Section) bool {
	return (other.start <= self.end && other.start >= self.start) || (other.end >= self.start && other.end <= self.end)
}

type Pair struct {
	FirstSection  Section
	SecondSection Section
}

func (self *Pair) HasCompleteOverlap() bool {
	return self.FirstSection.HasCompleteOverlapWith(self.SecondSection) || self.SecondSection.HasCompleteOverlapWith(self.FirstSection)
}

func (self *Pair) HasPartialOverlap() bool {
	return self.FirstSection.HasPartialOverlapWith(self.SecondSection) || self.SecondSection.HasPartialOverlapWith(self.FirstSection)
}

func CreateSectionFromString(line string) Section {
	parts := strings.Split(line, "-")
	start, err := strconv.ParseInt(parts[0], 10, 0)
	if err != nil {
		log.Printf("Could not parse start [input: %v]", parts[0])
	}
	end, err := strconv.ParseInt(parts[1], 10, 0)
	if err != nil {
		log.Printf("Could not parse end [input: %v]", parts[1])
	}

	return Section{
		start: start,
		end:   end,
	}
}

func CreatePairs() []Pair {
	var pairs []Pair

	onLineRead := func(line string) {
		parts := strings.Split(line, ",")
		pair := Pair{
			FirstSection:  CreateSectionFromString(parts[0]),
			SecondSection: CreateSectionFromString(parts[1]),
		}
		pairs = append(pairs, pair)
	}

	ReadFile("./resources/day_four_input.txt", onLineRead)

	return pairs
}

func DayFourA() int64 {
	pairs := CreatePairs()

	var pairsWithCompleteOverlapCount int64

	for _, pair := range pairs {
		if pair.HasCompleteOverlap() {
			pairsWithCompleteOverlapCount = pairsWithCompleteOverlapCount + 1
		}
	}

	return pairsWithCompleteOverlapCount
}

func DayFourB() int64 {
	pairs := CreatePairs()

	var pairsWithPartialOverlapCount int64

	for _, pair := range pairs {
		if pair.HasPartialOverlap() {
			pairsWithPartialOverlapCount = pairsWithPartialOverlapCount + 1
		}
	}

	return pairsWithPartialOverlapCount
}
func DayFour() {
	log.Print("Day 4")
	dayFourAResult := DayFourA()
	log.Printf("Result A: %v", dayFourAResult)
	dayFourBResult := DayFourB()
	log.Printf("Result B: %v", dayFourBResult)
}
