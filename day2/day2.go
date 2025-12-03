package day2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Part1() {
	f, _ := os.Open("data/day2/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	matchTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, line := range ranges {
			matchTotal += SumSymmetricNumbers(line)
		}
	}
	fmt.Printf("Total matches %v", matchTotal)
}

func Part2() {
	f, _ := os.Open("data/day2/input.txt")
	defer f.Close()

	scanner := bufio.NewScanner(f)
	matchTotal := 0
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, line := range ranges {
			matchTotal += SumRepeatNumbers(line)
		}
	}
	fmt.Printf("Total matches %v", matchTotal)
}

func SumSymmetricNumbers(line string) int {
	start, _ := strconv.Atoi(strings.Split(line, "-")[0])
	end, _ := strconv.Atoi(strings.Split(line, "-")[1])
	matchCount := 0

	for i := start; i <= end; i++ {
		number := strconv.Itoa(i)
		length := len(number)
		if length%2 == 1 {
			// Don't bother to check odd length numbers
			continue
		}

		left := number[:length/2]
		right := number[length/2:]

		if left == right {
			matchCount += i
		}
	}
	return matchCount
}

func SumRepeatNumbers(line string) int {
	start, _ := strconv.Atoi(strings.Split(line, "-")[0])
	end, _ := strconv.Atoi(strings.Split(line, "-")[1])
	matchCount := 0

	for i := start; i <= end; i++ {
		number := strconv.Itoa(i)
		length := len(number)

		// Check for repeat patterns, taking slices up to half the length of the number
		for j := 1; j <= length/2; j++ {
			// Don't bother checking if the size of the slice isn't a factor of the number's length
			if length%j != 0 {
				continue
			}
			slice := number[:j]
			repeat := strings.Repeat(slice, length/j)
			if repeat == number {
				matchCount += i
				break
			}

		}
	}
	return matchCount
}

func RunTests() {
	Test("11-22", 33, 33)
	Test("95-115", 99, 210)
	Test("998-1012", 1010, 2009)
	Test("1188511880-1188511890", 1188511885, 1188511885)
	Test("222220-222224", 222222, 222222)
	Test("1698522-1698528", 0, 0)
	Test("446443-446449", 446446, 446446)
	Test("38593856-38593862", 38593859, 38593859)
	Test("565653-565659", 0, 565656)
	Test("824824821-824824827", 0, 824824824)
	Test("2121212118-2121212124", 0, 2121212121)
}

func Test(line string, expected1 int, expected2 int) {
	actual := SumSymmetricNumbers(line)
	if expected1 != actual {
		fmt.Printf("Test error: line=%v expected1=%v actual=%v\n", line, expected1, actual)
	}

	actual = SumRepeatNumbers(line)
	if expected2 != actual {
		fmt.Printf("Test error: line=%v expected2=%v actual=%v\n", line, expected2, actual)
	}
}
