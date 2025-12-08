package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
)

func largestAfterRemovingDigits(s string, keep int) string {
	n := len(s)
	if n <= keep {
		return s // nie trzeba nic usuwać
	}
	toRemove := n - keep
	stack := []byte{}

	for i := 0; i < n; i++ {
		digit := s[i]

		// dopóki możemy usuwać, a na stosie jest mniejsza cyfra – wyrzucamy ją
		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < digit {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, digit)
	}

	// jeśli zostało coś do usunięcia (np. ciąg malejący) – usuwamy z końca
	if toRemove > 0 {
		stack = stack[:len(stack)-toRemove]
	}

	// wynik ma dokładnie 'keep' cyfr
	return string(stack[:keep])
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
	total := new(big.Int)
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		// wybieramy dokładnie 12 cyfr, tworząc największą możliwą liczbę
		resultStr := largestAfterRemovingDigits(line, 12)

		// konwertujemy na big.Int
		num := new(big.Int)
		num, ok := num.SetString(resultStr, 10)
		if !ok {
			panic("błąd konwersji: " + resultStr)
		}

		// dodajemy do sumy
		total.Add(total, num)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println("Całkowity output joltage:", total)
}
