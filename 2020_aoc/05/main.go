package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func find_mid(chr string, start *int, end *int) {
	mid := (*start + *end) / 2
	if chr == "F" || chr == "L" {
		*end = mid
	} else {
		*start = mid + 1
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
	seenSeats := make(map[int]bool)
	var ids_arr []int
	for scanner.Scan() {
		row_start, row_end := 0, 127
		col_start, col_end := 0, 7
		boardingPass := scanner.Text()
		for _, chr := range strings.Split(boardingPass, "") {
			if chr == "F" || chr == "B" {
				find_mid(chr, &row_start, &row_end)
			}
			if chr == "L" || chr == "R" {
				find_mid(chr, &col_start, &col_end)
			}
		}
		seat_id = row_end*8 + col_end
		ids_arr = append(ids_arr, seat_id)
		seenSeats[seat_id] = true
		if seat_id > highest_id {
			highest_id = seat_id
		}
	}
	sort.Ints(ids_arr)
	curr := 0
	for i := ids_arr[0]; i < ids_arr[len(ids_arr)-1]; i++ {
		if !seenSeats[i] && seenSeats[i+1] && seenSeats[i-1] {
			curr = i
			break
		}
	}
	fmt.Println(highest_id, curr)
}
