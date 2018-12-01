package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	sum := 0

	f, err := os.Open("input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		sum += num
	}
	fmt.Printf("%d\n", sum)
}

		
