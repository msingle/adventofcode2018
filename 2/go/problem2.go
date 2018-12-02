package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func similar(s, maybe string) ([]string, int) {
	similars := []string{}
	nomatches := 0
	initLetters := strings.Split(s, "")
	maybeLetters := strings.Split(maybe, "")
	for x := range  initLetters {
		if initLetters[x] == maybeLetters[x] {
			similars = append(similars, initLetters[x])
		} else {
			nomatches++
		}
	}
	return similars, nomatches
}
func main() {

	f, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	for i, line := range lines {
		for _, next := range lines[i:] {
			matches, misses := similar(line, next)
			if misses == 1 {
				fmt.Printf("%s", matches)
			}
		}
	}

}
