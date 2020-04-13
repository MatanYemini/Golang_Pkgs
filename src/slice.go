package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0, 3)

	for {
		var input string

		fmt.Println("Enter a number press x for exit: ")
		fmt.Scan(&input)

		input = strings.ToLower(input)
		if input == "x" {
			fmt.Println("Bye!")
			os.Exit(0)
		}

		num, err := strconv.Atoi(input)

		if err != nil {
			fmt.Println("Bad input")
			os.Exit(1)
		}

		s = append(s, num)
		sort.Slice(s, func(i, j int) bool {
			return s[i] < s[j]
		})
		fmt.Println(s)

	}

}
