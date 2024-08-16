package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var numbers []int
	number_map := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		text := scanner.Text()
		// if err != nil {
		// 	log.Fatal(err)
		// }
		num, _ := strconv.Atoi(text)
		if num < 2020 {
			reminder := 2020 - num
			if _, ok := number_map[strconv.Itoa(reminder)]; ok {
				fmt.Println(reminder * num)
				// os.Exit(0)
			} else {
				number_map[strconv.Itoa(num)] = reminder
			}
		}
		numbers = append(numbers, num)
	}
	fmt.Println(numbers)
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			for k := j + 1; k < len(numbers); k++ {
				if numbers[i]+numbers[j]+numbers[k] == 2020 {
					fmt.Println(numbers[i], numbers[j], numbers[k])
					fmt.Println(numbers[i] * numbers[j] * numbers[k])
				}
			}
		}

	}
}
