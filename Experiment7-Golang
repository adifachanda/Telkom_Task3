package main

import "fmt"

func logPairs(array1 []string, array2 []string, words string) {
	counter := 0
	for i := 0; i < len(array1); i++ {
		for j := 0; j < len(array2); j++ {
			counter++
			fmt.Println(words, array1[i], "dan", array2[j])
		}
	}
}

func main() {
	foods := []string{"sate", "bakso", "dimsum", "rames"}
	drinks := []string{"jeruk", "teh", "kelapa", "cendol"}
	logPairs(foods, drinks, "Menu")
}
