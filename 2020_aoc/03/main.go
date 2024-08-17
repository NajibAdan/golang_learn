package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func count_trees(step_x int, step_y int) int {
	// open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// declare variable to move right
	x := 0
	num_of_trees := 0

	// initialize the starting position
	if step_y == 1 {
		for i := 0; i < step_y; i++ {
			scanner.Scan()
		}
	} else {
		for i := 1; i < step_y; i++ {
			scanner.Scan()
		}
	}
	for scanner.Scan() {
		// go y down
		for i := 1; i < step_y; i++ {
			scanner.Scan()
		}
		// go x right
		x += step_x
		if x >= 31 {
			// if go the limits we reset ourself
			x = x - 31
		}
		line := []rune(scanner.Text())
		if line[x] == '#' {
			// tree was found
			num_of_trees++
		}
	}
	return num_of_trees
}

func main() {
	// steps to move across
	step_x := [...]int{1, 3, 5, 7, 1}
	step_y := [...]int{1, 1, 1, 1, 2}
	res := 1
	for i := 0; i < len(step_x); i++ {
		trees := count_trees(step_x[i], step_y[i])
		fmt.Println(step_x[i], trees)
		res *= trees
	}
	fmt.Println(res)

}
