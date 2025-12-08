package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"path/filepath"
)


// Finding largest int with index
func findLargest(batteries string) (int, int) {
	max := 0
	maxIndex := 0 
	arr := strings.Split(batteries, "")
	for i := 0; i < len(arr); i++ {
		number, err := strconv.Atoi(arr[i])
		if err != nil {
			panic(err)
		}
		if number > max {
			max = number
			maxIndex = i
		}
	}
	return max, maxIndex
}

func main() {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os_path := filepath.Dir(ex)
	path := filepath.Join(os_path, "aoc_solutions/cmd/day3_input.txt")
	_, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loading input from: %s", string(path))
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	sum := 0

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		max, maxIndex := findLargest(line[0:len(line)-1])
		maxSecond, maxindex2 := findLargest(line[maxIndex : ])
		jotlage := fmt.Sprintf("%d%d", max, maxSecond)
		jolatge_num, err := strconv.Atoi(jotlage)
		if err != nil {
			panic(err)
		}
		sum += jolatge_num
		fmt.Print(sum, maxindex2)
	}
}
