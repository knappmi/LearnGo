package main

import (
	"fmt"
	"strconv"
)

/*
Problem 2 — FizzBuzz (Loops & Conditionals)

Goal:

	Print numbers from 1 to 100. Replace multiples of 3 with "Fizz", 5 with "Buzz",
	and multiples of both with "FizzBuzz".

Why:

	Teaches for-loops, if/else, and modulus arithmetic.

Run:

	go run ./2_fizzbuzz

Expected (start):

	1
	2
	Fizz
	4
	Buzz
	Fizz
	7
	8
	Fizz
	Buzz
	11
	Fizz
	13
	14
	FizzBuzz
	...
*/
func main() {
	var str string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			str = "FizzBuzz"
		} else if i%3 == 0 {
			str = "Fizz"
		} else if i%5 == 0 {
			str = "Buzz"
		} else {
			str = strconv.Itoa(i)
		}
		fmt.Println(str)
	}
}
