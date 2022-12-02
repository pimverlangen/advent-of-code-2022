package main

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(path string, onLineRead func(string)) {
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		var line = scanner.Text()
		onLineRead(line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
