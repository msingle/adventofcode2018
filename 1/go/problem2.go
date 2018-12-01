package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	curr := 0

	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	seen := map[int]bool{}
	changes := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		changes = append(changes, num)
	}
	repeat := true
	for repeat {
		for _, change := range changes {
			curr += change
			if _, ok := seen[curr]; ok {
				repeat = false
				break
			} else {
				seen[curr] = true
			}
		}
	}
	fmt.Println(curr)
}

		
