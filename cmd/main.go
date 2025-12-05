package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"path/filepath"
)

func isInvalid(s string) bool {
	n := len(s)
	for l := 1; l <= n/2; l++ {
		if n%l != 0 {
			continue
		}
		k := n / l
		if k < 2 {
			continue
		}
		substr := s[:l]
		valid := true
		for i := 1; i < k; i++ {
			if s[i*l:(i+1)*l] != substr {
				valid = false
				break
			}
		}
		if valid {
			return true
		}
	}
	return false
}

func main() {
ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os_path := filepath.Dir(ex)
	path := filepath.Join(os_path, "aoc_solutions/cmd/day2_input.txt")
	_, err = os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Loading input from: %s", string(path))
	f, err := os.Open(path)

	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	scanner.Scan()
	line := scanner.Text()
	ranges := strings.Split(line, ",")

	// Fix: Filter out empty ranges (e.g., from trailing comma)
	var filteredRanges []string
	for _, r := range ranges {
		if r != "" {
			filteredRanges = append(filteredRanges, r)
		}
	}
	ranges = filteredRanges

	var total uint64 = 0
	for _, r := range ranges {
		parts := strings.Split(r, "-")
		if len(parts) != 2 {  // Additional safety: Skip malformed ranges
			continue
		}
		start, err1 := strconv.ParseUint(parts[0], 10, 64)
		if err1 != nil {
			continue
		}
		end, err2 := strconv.ParseUint(parts[1], 10, 64)
		if err2 != nil || start > end {  // Also check start <= end
			continue
		}
		for num := start; num <= end; num++ {
			s := strconv.FormatUint(num, 10)
			fmt.Println(s)
			if isInvalid(s) {
				total += num
			}
		}
	}
	fmt.Println(total)
}
