package day1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Dial struct {
	pos       int
	zeroCount int
}

func (d *Dial) Turn(line string) {
	number, err := strconv.Atoi(line[1:])
	if err != nil {
		log.Fatal(line + " cannot be parsed as number")
	}

	d.zeroCount += number / 100
	number = number % 100

	if line[0] == 'L' {
		if d.pos == number {
			// end on zero
			if d.pos != 0 {
				// only count if we didn't start on zero
				d.zeroCount += 1
			}
			d.pos = 0
		} else if d.pos < number {
			// go beyond zero
			if d.pos != 0 {
				// only count if we didn't start on zero
				d.zeroCount += 1
			}
			d.pos = d.pos - number + 100
		} else if d.pos > number {
			// don't get to zero
			d.pos = d.pos - number
		}
	} else {
		// Must be R
		if d.pos+number == 100 {
			// end on zero
			d.pos = 0
			d.zeroCount += 1
		} else if d.pos+number > 100 {
			// go beyond zero
			d.pos = d.pos + number - 100
			d.zeroCount += 1
		} else if d.pos+number < 100 {
			// don't get to zero
			d.pos = d.pos + number
		}
	}

}

func RunTests() {
	Test(50, "L49", 0)
	Test(50, "L50", 1)
	Test(50, "L51", 1)
	Test(50, "L100", 1)
	Test(50, "L150", 2)
	Test(50, "L250", 3)
	Test(50, "R49", 0)
	Test(50, "R50", 1)
	Test(50, "R51", 1)
	Test(50, "R100", 1)
	Test(50, "R150", 2)
	Test(50, "R250", 3)
	Test(0, "L50", 0)
	Test(0, "L100", 1)
	Test(0, "L150", 1)
	Test(0, "L200", 2)
}

func Test(start int, line string, expected int) {
	d := Dial{
		pos:       start,
		zeroCount: 0,
	}
	d.Turn(line)
	if expected != d.zeroCount {
		fmt.Printf("Test error: start=%v turn=%v expected=%v actual=%v\n", start, line, expected, d.zeroCount)
	}
}

func Part1() {
	pos := 50
	zeroCount := 0

	f, err := os.Open("data/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(line + " cannot be parsed as number")
		}

		number = number % 100

		if line[0] == 'L' {
			pos = pos - number
			if pos < 0 {
				pos = pos + 100
			}
		} else {
			// Must be R
			pos = pos + number
			if pos > 99 {
				pos = pos - 100
			}
		}

		if pos == 0 {
			zeroCount += 1
		}

		fmt.Printf(line+" position now %d\n", pos)
	}
	fmt.Printf("Finished with %v zeros\n", zeroCount)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

}

func Part2() {
	f, err := os.Open("data/day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	d := Dial{
		pos:       50,
		zeroCount: 0,
	}
	for scanner.Scan() {
		line := scanner.Text()
		d.Turn(line)
	}
	fmt.Printf("Finished with %v zeros\n", d.zeroCount)
	// 6513 too high

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
