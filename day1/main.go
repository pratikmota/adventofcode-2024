package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part2(left []int, right []int) int {
	total := 0
	for _, v := range left {
		// check v number available in right how many time?
		count := 0
		for _, rv := range right {
			if v == rv {
				count++
			}
		}
		total += (v * count)
	}

	return total
}

func part1(left []int, right []int) int {

	sort.Ints(left)
	sort.Ints(right)
	total := 0
	for i := 0; i < len(left); i++ {
		total += int(math.Abs(float64(left[i] - right[i])))
	}

	return total
}

func main() {

	left, right, err := readValuesFromFile("input.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	fmt.Println(part1(left, right))
	fmt.Println(part2(left, right))
}

func readValuesFromFile(filename string) ([]int, []int, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left, right []int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		if len(values) == 2 {
			leftValue, _ := strconv.Atoi(values[0])
			rightValue, _ := strconv.Atoi(values[1])
			left = append(left, leftValue)
			right = append(right, rightValue)
		}
	}

	return left, right, scanner.Err()
}
