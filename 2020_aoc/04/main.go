package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func check_valid_passport_part_two(features map[string]string) bool {
	byr, byr_ok := features["byr"]
	iyr, iyr_ok := features["iyr"]
	eyr, eyr_ok := features["eyr"]
	hgt, hgt_ok := features["hgt"]
	hcl, hcl_ok := features["hcl"]
	ecl, ecl_ok := features["ecl"]
	pid, pid_ok := features["pid"]
	if !byr_ok || !iyr_ok || !eyr_ok || !hgt_ok || !hcl_ok || !ecl_ok || !pid_ok {
		return false
	}
	if check_yr(byr, 1920, 2002) && check_yr(iyr, 2010, 2020) && check_yr(eyr, 2020, 2030) && check_hgt(hgt) && check_hcl(hcl) && check_ecl(ecl) && check_pid(pid) {
		return true
	}
	return false
}

func check_ecl(ecl string) bool {
	if ecl == "amb" || ecl == "blu" || ecl == "brn" || ecl == "gry" || ecl == "grn" || ecl == "hzl" || ecl == "oth" {
		return true
	}
	return false
}

func check_pid(pid string) bool {
	m, err := regexp.MatchString(`\b\d{9}\b`, pid)
	if err != nil {
		return false
	}
	if m {
		return true
	}
	return false
}

func check_hcl(hcl string) bool {
	m, err := regexp.MatchString(`\#[0-9a-fA-F]{6}`, hcl)
	if err != nil {
		return false
	}
	if m {
		return true
	}
	return false
}

func check_hgt(hgt string) bool {
	if len(hgt) == 5 {
		// height in cm
		hgt_c, err := strconv.Atoi(strings.Split(hgt, "cm")[0])
		if err != nil {
			return false
		}
		if hgt_c >= 150 && hgt_c <= 193 {
			return true
		}
	}
	if len(hgt) == 4 {
		// height in inches
		hgt_c, err := strconv.Atoi(strings.Split(hgt, "in")[0])
		if err != nil {
			return false
		}
		if hgt_c >= 59 && hgt_c <= 76 {
			return true
		}
	}
	return false
}

func check_yr(iyr string, min int, max int) bool {
	if len(iyr) < 4 {
		return false
	}
	yr, err := strconv.Atoi(iyr)
	if err != nil {
		return false
	}
	if yr >= min && yr <= max {
		return true
	}
	return false
}

func check_valid_passport(features map[string]string) bool {
	if len(features) == 8 {
		return true
	}
	if len(features) <= 6 {
		return false
	}
	// _, pid_ok := features["pid"]
	_, cid_ok := features["cid"]
	if !cid_ok {
		return true
	}
	return false
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	features := make(map[string]string)
	correct_passports := 0
	part_two_passports := 0
	var line_str []string
	for scanner.Scan() {
		line := scanner.Text()
		// fmt.Println(line)
		if line == "" {
			for _, feature := range line_str {
				for _, feat := range strings.Split(feature, " ") {
					str_arr := strings.Split(feat, ":")
					features[str_arr[0]] = str_arr[1]
				}
			}
			if check_valid_passport(features) {
				correct_passports++
			}
			if check_valid_passport_part_two(features) {
				part_two_passports++
			}
			line_str = nil

			features = make(map[string]string)
			// break
		} else {
			line_str = append(line_str, line)
		}
	}
	if scanner.Text() == "" {
		for _, feature := range line_str {
			for _, feat := range strings.Split(feature, " ") {
				str_arr := strings.Split(feat, ":")
				features[str_arr[0]] = str_arr[1]
			}
		}
		if check_valid_passport(features) {
			correct_passports++
		}
		if check_valid_passport_part_two(features) {
			part_two_passports++
		}
	}
	fmt.Println(correct_passports, part_two_passports)

}
