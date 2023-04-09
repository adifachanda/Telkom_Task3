package main

import (
	"fmt"
	"time"
)

const company string = "telkom"

func findCompany(array [10]string) {
	tx := time.Now()
	for i := 0; i < len(array); i++ {
		if array[i] == "telkom" {
			fmt.Println("Company Found!")
		}
	}
	ty := time.Since(tx)
	fmt.Println("The Performance is", ty.Milliseconds(), "ms")
}

func main() {
	var largeCompanyName [10]string
	for i := 0; i < len(largeCompanyName); i++ {
		largeCompanyName[i] = company  
	}
	findCompany(largeCompanyName)
}
