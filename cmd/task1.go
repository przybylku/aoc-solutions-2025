package main
// to cleanup
// Author Przybylku
import (
	"bufio"
	"fmt"
	"strconv"
	"os"
	"path/filepath"
)

// AOC solutions 2025 by Przybylku



// Gets start points, direction and points to rotate, returns new point.
func rotate(start int, direction int, points int) (int, int) {
	_p := start 
	pointed_zeros := 0
	// rotating with checking if next point is out of index.
	for i := 0; i < points; i++ {
		if _p == 0 && i != 0{ 
			pointed_zeros++
		}
		if _p == 99 && direction == 82 {
			_p = 0
		} else if _p == 0 && direction == 76 {
			_p = 99 
		}else if direction == 82 {
			_p++
		} else if direction == 76 {
			_p--
		}
		
	}
	return _p, pointed_zeros
}

func main() {
	ex, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	os_path := filepath.Dir(ex)
	path := filepath.Join(os_path, "aoc_solutions/cmd/input.txt")
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
	pos := 50
	code := 0
	for scanner.Scan() { 
	    line := scanner.Text()
        dir := line[0]
	    points, _ := strconv.Atoi(line[1:])
		rotated, pointed_zeros := rotate(pos, int(dir), points)
		pos = rotated
		if rotated == 0 {
			code++
		}
		code += pointed_zeros
		fmt.Printf("The dial is rotated %s to point at %d. During this rotation, it points 0 - %d times, a code is now: %d \n", line, rotated, pointed_zeros, code)
	}
	fmt.Println("The code: ", code)
}

