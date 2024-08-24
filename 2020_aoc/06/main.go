package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func processGroup(questions *[]string) (int, int) {
	groupAnswered := make(map[rune]int)
	for _, question := range *questions {
		for _, chr := range question {
			groupAnswered[chr]++
		}
	}
	allYesCount := 0
	for _, count := range groupAnswered {
		if count == len(*questions) {
			allYesCount++
		}
	}
	return len(groupAnswered), allYesCount
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var questions []string
	total_sum_1 := 0
	total_sum_2 := 0
	for scanner.Scan() {
		line_text := scanner.Text()
		if line_text == "" {
			sum1, sum2 := processGroup(&questions)
			total_sum_1 += sum1
			total_sum_2 += sum2
			questions = nil
		} else {

			questions = append(questions, line_text)
		}
	}

	sum1, sum2 := processGroup(&questions)
	total_sum_1 += sum1
	total_sum_2 += sum2
	fmt.Println(total_sum_1, total_sum_2)
}
