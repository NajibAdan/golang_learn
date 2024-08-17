package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// Counts the number of times a letter appears in the word
func check_count_of_character(word string, letter string) int {
	curr_count := 0
	for _, char := range word {
		if string(char) == letter {
			curr_count++
		}
	}
	return curr_count
}

// checks if the password is valid or not based on the shopkeepers old job
func check_password(minimum int, maximum int, letter_to_contain string, password string) bool {
	letter_to_contain_arr := []rune(letter_to_contain)
	letter_rune := string(letter_to_contain_arr[0])

	chr_count := check_count_of_character(password, letter_rune)
	if minimum <= chr_count && chr_count <= maximum {
		return true
	} else {
		return false
	}
}

// checks if the password is valid or not based on the Toboggan Corporate Policy
func toboggan_password_checker(start int, end int, letter string, password string) bool {
	password_arr := []rune(password)
	letter_rune := []rune(letter)[0]
	if password_arr[start] == letter_rune && password_arr[end] != letter_rune {
		return true
	} else if password_arr[start] != letter_rune && password_arr[end] == letter_rune {
		return true
	}
	return false
}

func main() {
	// read the file
	input_file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("File not found")
	}
	defer input_file.Close()
	scanner := bufio.NewScanner(input_file)

	// variables to hold if the password or not based on the old password policy and the official one
	count_old_correct_passwords := 0
	count_correct_toboggan_passwords := 0

	// loop through line by line
	for scanner.Scan() {
		// parse the line
		split_string := strings.Split(scanner.Text(), " ")

		// extract the numbers from the first section
		limits := strings.Split(split_string[0], "-")
		minimum, _ := strconv.Atoi(limits[0])
		maximum, _ := strconv.Atoi(limits[1])

		// check if the password is correct based on the old password policy
		if check_password(minimum, maximum, split_string[1], split_string[2]) {
			count_old_correct_passwords++
		}
		// check if the password is correct based on the official toboggan policy
		if toboggan_password_checker(minimum-1, maximum-1, split_string[1], split_string[2]) {
			count_correct_toboggan_passwords++
		}
	}
	fmt.Println(count_old_correct_passwords, count_correct_toboggan_passwords)
}
