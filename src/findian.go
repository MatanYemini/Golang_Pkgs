package main

import (
	"fmt"
	"os"
	s "strings"
)

func main() {
	var input_str string

	fmt.Printf("Enter a string: ")
	_, err := fmt.Scanf("%s", &input_str)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	input_str = s.ToLower(input_str)
	if s.HasPrefix(input_str, "i") &&
		s.HasSuffix(input_str, "n") &&
		s.Contains(input_str, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
