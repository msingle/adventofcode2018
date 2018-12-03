package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type plan struct {
	id int
	startPoint point
	width, height int
}

func parseId(s string) int {
	chunks := strings.Split(s, "#")
	id, err := strconv.Atoi(chunks[1])
	if err != nil {
		panic(err)
	}
	return id
}

func parseArea(s string) (int, int) {
	chunks := strings.Split(s, "x")
	width, err := strconv.Atoi(chunks[0])
	if err != nil {
		panic(err)
	}
	height, err := strconv.Atoi(chunks[1])
	if err != nil {
		panic(err)
	}
	return width, height
}

func parsePoint(s string) point {
	chunks := strings.Split(s, ",")
	x, err := strconv.Atoi(chunks[0])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strings.TrimRight(chunks[1], ":"))
	if err != nil {
		panic(err)
	}
	return point{ x:x, y:y}
}

func newPlan(line string) plan {
	fields := strings.Fields(line)
	width, height := parseArea(fields[3])
	return plan{
		id: parseId(fields[0]),
		startPoint: parsePoint(fields[2]),
		width: width,
		height: height,
	}
}

func (p plan) endX() int {
	return p.startPoint.x	+ p.width
}

func (p plan) endY() int {
	return p.startPoint.y	+ p.height
}
func addPlansToTally(tally map[point][]plan, p plan) {
	x := p.startPoint.x + p.width
	y := p.startPoint.y + p.height
	for i := p.startPoint.x; i < x; i++ {
		for j := p.startPoint.y; j < y; j++ {
			pt := point{x:i, y:j}
			//fmt.Printf("adding %v\n", pt)
			tally[pt] = append(tally[pt], p)
		}
	}
}
func main() {

	f, err := os.Open("../input")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer f.Close()

	plans := map[int]plan{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		plan := newPlan(line)
		plans[plan.id] = plan
	}

	tally := map[point][]plan{}
	for _, plan := range plans {
		//fmt.Printf("adding points from %v\n", plan)
		addPlansToTally(tally, plan)
	}
	seen := map[int]plan{}
	for _, plans := range tally {
		if len(plans) > 1 {
			for _, plan := range plans {
				seen[plan.id] = plan
			}
		}
	}
	for _, plan := range plans {
		if _, ok := seen[plan.id]; !ok {
			fmt.Printf("%v\n", plan)
		}
	}
}
