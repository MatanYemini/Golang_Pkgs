package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Name struct {
	_fname string
	_lname string
}

func main() {
	var namesSlice []Name
	var input string

	fmt.Println("Enter a file to read from: ")
	fmt.Scan(&input)
	file, err := os.Open(input)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close() // Defer determine that it should execute only when the function ends

	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		lineText := string(line)
		s := strings.Split(lineText, " ")
		var full_name Name
		full_name._fname = s[0]
		full_name._lname = s[1]
		namesSlice = append(namesSlice, full_name)
	}

	// Now printing
	for _, n := range namesSlice {
		fmt.Printf("%s %s", n._fname, n._lname)
	}
}
