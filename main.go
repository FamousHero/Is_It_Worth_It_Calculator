package main

import (
	"fmt"
	"os"
	"strconv"
)

type PayInfo struct {
	BasePay int
	Hours   int
}

func (pi PayInfo) String() string {
	return fmt.Sprintf("$%d, %dhrs", pi.BasePay, pi.Hours)
}

func parseArgs(args []string) []int {
	var parsed_args []int
	for _, arg := range args {
		i, err := strconv.Atoi(arg)
		if err != nil {
			// Create better error message
			panic(err)
		}
		parsed_args = append(parsed_args, i)
	}
	return parsed_args
}

func main() {
	/*
		take in 4 arguments: pay_1, hours_1, pay_2, hours_2
		Break down into primes & store in a set
		reduce common primes
		multiply outliers for each separately and compare
		return j1 > j2? j1 : j2
	*/

	args := os.Args[1:]
	p_args := parseArgs(args)
	if len(p_args) != 4 {
		panic("Please enter 4 numbers in the order: pay1, hours1, pay2, hours2")
	}
	j1 := PayInfo{
		BasePay: p_args[0],
		Hours:   p_args[1],
	}
	j2 := PayInfo{
		BasePay: p_args[2],
		Hours:   p_args[3],
	}

	fmt.Printf("Job 1: %v\nJob 2: %v", j1, j2)
}
