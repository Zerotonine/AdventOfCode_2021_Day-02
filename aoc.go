package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func getInput() []string {
	file, err := os.Open("input.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("Failed to open!")
	}
	scanner := bufio.NewScanner(file)
	var text []string
	for scanner.Scan() {
		text = append(text, scanner.Text())
	}
	return text
}

func solutionOne(text *[]string) int {
	var horizontal, depth int = 0, 0
	for _, line := range *text {
		direction := strings.Fields(line)[0]
		steps, _ := strconv.Atoi(strings.Fields(line)[1])

		switch direction {
		case "forward":
			horizontal += steps
		case "down":
			depth += steps
		case "up":
			depth -= steps
		}
	}
	return horizontal * depth
}

func solutionTwo(text *[]string) int {
	var horizontal, depth, aim int = 0, 0, 0
	for _, line := range *text {
		direction := strings.Fields(line)[0]
		steps, _ := strconv.Atoi(strings.Fields(line)[1])

		switch direction {
		case "forward":
			horizontal += steps
			depth += aim * steps
		case "down":
			aim += steps
		case "up":
			aim -= steps
		}
	}
	return horizontal * depth
}

func main() {
	text := getInput()
	fmt.Println("Answer 1:", solutionOne(&text))
	fmt.Println("Answer 2:", solutionTwo(&text))
}
