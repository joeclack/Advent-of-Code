package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	total := 0

	data, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	values := strings.Split(strings.TrimSpace(string(data)), ",")

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

	fmt.Println(total)
}
