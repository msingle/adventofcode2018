package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func repetitionsIn (s string, repetitions int) int {
	tally := map[string]int{}
	for _, letter := range strings.Split(s,"") {
		tally[letter]++
	}
	fmt.Println(tally)
	count := 0
	for _, v := range tally {
		if v == repetitions {
			count = 1
		}
	}
	return count
}

func main() {

	f, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	doubles := 0
	triples := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		doubles += repetitionsIn(line, 2)
		triples += repetitionsIn(line, 3)
	}

	fmt.Printf("%d * %d = %d", doubles, triples, doubles * triples)
}
