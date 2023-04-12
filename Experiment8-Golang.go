package main

import "fmt"

var results [][]string
var candidates = []string{"Baswedan", "Subianto", "Maharani"}

func arrange(array []string, memory []string) [][]string {
	var current string
	if memory == nil {
		memory = []string{}
	}

	for i := 0; i < len(array); i++ {
		current = array[i]
		array = append(array[:i], array[i+1:]...)

		if len(array) == 0 {
			results = append(results, append(memory, current))
		}

		arrange(append([]string{}, array...), append(memory, current))
		array = append(array[:i], append([]string{current}, array[i:]...)...)
	}
	return results
}

func main() {
	chairs := arrange(candidates, nil)
	for _, c := range chairs {
		fmt.Println(c)
	}
}
