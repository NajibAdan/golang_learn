package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
)

func find_mid(chr string, start *int, end *int) {
	if chr == "F" || chr == "L" {
		*end = (*start + *end) / 2
	} else {
		*start = (*start + *end) / 2
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	seat_id := 0
	highest_id := -1
	var ids_arr []int
	for scanner.Scan() {
		row_start, row_end := 0, 127
		col_start, col_end := 0, 7
		arrangement := scanner.Text()
		for _, chr := range strings.Split(arrangement, "") {
			if chr == "F" || chr == "B" {
				find_mid(chr, &row_start, &row_end)
			}
			if chr == "L" || chr == "R" {
				find_mid(chr, &col_start, &col_end)
			}
		}
		seat_id = row_end*8 + col_end
		ids_arr = append(ids_arr, seat_id)
		if seat_id > highest_id {
			highest_id = seat_id
		}
	}
	slices.Sort(ids_arr)
	curr := ids_arr[0]
	for i := 1; i < len(ids_arr); i++ {
		curr++
		if ids_arr[i] != curr {
			break
		}
	}
	fmt.Println(highest_id, curr)
}
