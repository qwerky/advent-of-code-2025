package day3

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	f, _ := os.Open("data/day3/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalJoltage := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalJoltage += FindMaxJoltage(line)

	}
	fmt.Printf("Max joltage %v", totalJoltage)
}

func Part2() {
	f, _ := os.Open("data/day3/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	totalJoltage := 0
	for scanner.Scan() {
		line := scanner.Text()
		totalJoltage += FindMegaJoltage(line)

	}
	fmt.Printf("Mega joltage %v", totalJoltage)
}

func FindMaxJoltage(line string) int {
	// Lets just brute force it, starting with the highest and
	// just try to find if it's there with a regex
	for tens := 9; tens >= 0; tens-- {
		for units := 9; units >= 0; units-- {
			regex := regexp.MustCompile(fmt.Sprintf("%v.*%v", tens, units))
			if regex.MatchString(line) {
				return (tens * 10) + units
			}
		}
	}
	panic("low joltage!")
}

func FindMegaJoltage(line string) int {
	// Don't think we can brute force 12 digit numbers as regex :(

	// Strategy is to start scanning for the largest, leftmost digits one at a time
	// (leftmost 9, then leftmost 8, etc). Once the largest, leftmost digit is found
	// we start again with the next digit, but start scanning from the next character.
	// start variable holds
	// end index holds

	// the start index of our scan. During the loop when we find a high digit, we
	// update this to point to the digit one position to the right of the found digit
	start := 0

	// the last index where the digit can be, increases by 1 each time we find a
	// digit - we can't scan all the way to the end when there are digits we must be
	// able to fit in afterwards
	end := len(line) - 11

	joltage := []string{}

	for len(joltage) < 12 {
		for n := 9; n > 0; n-- {
			// Check to see if n is in a substring start:end
			sub := line[start:end]
			index := strings.Index(sub, strconv.Itoa(n))
			if index != -1 {
				// We found the next highest digit!
				// Add it to the joltage and update the start and end,
				// then move on to the next digit
				joltage = append(joltage, strconv.Itoa(n))
				start = start + index + 1
				end = min(end+1, len(line))
				break
			}
		}
	}

	result, err := strconv.Atoi(strings.Join(joltage, ""))
	if err != nil {
		panic(fmt.Sprintf("Joltage not a number %v", joltage))
	}
	return result
}

func RunTests() {
	Test("987654321111111", 98, 987654321111)
	Test("811111111111119", 89, 811111111119)
	Test("234234234234278", 78, 434234234278)
	Test("818181911112111", 92, 888911112111)
}

func Test(line string, expected1 int, expected2 int) {
	actual := FindMaxJoltage(line)
	if expected1 != actual {
		fmt.Printf("Test error: line=%v expected1=%v actual=%v\n", line, expected1, actual)
	}

	actual = FindMegaJoltage(line)
	if expected2 != actual {
		fmt.Printf("Test error: line=%v expected2=%v actual=%v\n", line, expected2, actual)
	}
}
