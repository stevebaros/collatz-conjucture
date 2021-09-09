package main

import (
	"flag"
	"fmt"
	"math"
)

func even(n float64) (b bool) {
	if math.Mod(n, 2) == 0 {
		b = true
	} else {
		b = false
	}
	return
}

func odd(n float64) (b bool) {
	if even(n) {
		b = false
	} else {
		b = true
	}
	return
}

func collatz(n float64) {

	switch {
	case n == 1:
		fmt.Println("Reached the bottom!")
	case even(n):
		n = n / 2
		fmt.Println(n)
		collatz(n)
	case odd(n):
		n = n*3 + 1
		fmt.Println(n)
		collatz(n)
	}
}

func main() {
	flag.Parse()
	collatz(27)
}