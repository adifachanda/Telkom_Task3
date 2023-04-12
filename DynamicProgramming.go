package main

import "fmt"

var counterPlainRecursive int
var counterDynamicProgramming int

func fibonacciPlainRecursive(n int) int {
	counterPlainRecursive++
	if n < 2 {
		return n
	}
	return fibonacciPlainRecursive(n-1) + fibonacciPlainRecursive(n-2)
}

func fibonacciDynamicProgramming() func(int) int {
	var cache = make(map[int]int)
	var fib func(n int) int
	fib = func(n int) int {
		counterDynamicProgramming++
		if val, ok := cache[n]; ok {
			return val
		} else {
			if n < 2 {
				return n
			} else {
				val := fib(n-1) + fib(n-2)
				cache[n] = val
				return val
			}
		}
	}
	return fib
}

func main() {
	fibonacciPlainRecursive(20)
	fasterFibonacci := fibonacciDynamicProgramming()
	fasterFibonacci(20)
	fmt.Printf("we did %v calculations for Plain Recursive\n", counterPlainRecursive)
	fmt.Printf("we did %v calculations for Dynamic Programming\n", counterDynamicProgramming)
}
