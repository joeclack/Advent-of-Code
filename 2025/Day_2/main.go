package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	values := strings.Split(strings.TrimSpace(string(data)), ",")

	// part 1 - invalid id is string repeated i.e 1212 would be 12 repeated twice or 134134 is 134 repeated twice

	total := 0

	for _, v := range values {
		sten := strings.Split(v, "-")
		start, err := strconv.Atoi(sten[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(sten[1])
		if err != nil {
			panic(err)
		}

		valid := 0

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			le := len(str)
			if le%2 == 0 {
				mid := le / 2
				left := str[:mid]
				right := str[mid:]
				if left == right {
					valid += i
				}
			}
		}

		total += valid
	}

	out := fmt.Sprintf("Part 1: %d", total)
	fmt.Println(out)

	// part 2 - invalid ids are now strings with repeated sequences of digits at leat twice
	// i.e 121212 is 12 repeated 3 times or 1111 is 1 repeated 4 times

	total = 0 // reset total

	for _, v := range values {
		sten := strings.Split(v, "-")
		start, err := strconv.Atoi(sten[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(sten[1])
		if err != nil {
			panic(err)
		}

		valid := 0
		count := 0

		for i := start; i <= end; i++ {
			str := strconv.Itoa(i)
			le := len(str)
			max := le / 2
			for pat := 1; pat <= max; pat++ {
				if le%pat != 0 {
					continue
				}

				pattern := str[:pat]
				repeatCount := le / pat

				repeated := strings.Repeat(pattern, repeatCount)

				if repeated == str {
					if repeatCount >= 2 {
						valid += i
						count++
						pat = max + 1 // now this number has been identified, it doesnt need checking again.
						// 111111 would be flagged as invalid 3 times cause 1 repeated 6 times, 11 repeated 3 times and 111 repeated 2 times.
						// this should only be glagged once so move onto the next number in the range
					}
				}
			}
		}

		total += valid
	}

	out2 := fmt.Sprintf("Part 2: %d", total)
	fmt.Println(out2)
}
