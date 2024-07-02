package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func find_nums_in_string(str string) string {
	var s string
	for _, c := range str {
		switch c {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			s += string(c)
		}
	}
	return s
}

func str_to_num(str string) int {
	num, _ := strconv.Atoi(str)
	return num
}

func solver(str string) string {
	s := find_nums_in_string(str)
	if len(s) > 2 {
		var t string = string(s[0]) + string(s[len(s)-1])
		return t
	} else if len(s) < 2 {
		return s + s
	}
	return s
}

func main() {
	bytes, err := os.ReadFile("./ex.txt")
	if err != nil {
		panic("ur mom")
	}
	input := strings.Split(string(bytes), "\n")
	var output []string
	for _, s := range input {
		output = append(output, find_nums_in_string(s))

		/*
			for _, c := range s {
				switch c {
				case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
					output = append(output, string(c))
				}
			}
		*/
	}
	var num int
	for i := range output {
		p, _ := strconv.Atoi(solver(output[i]))
		num += p
	}
	fmt.Println(num)
}
