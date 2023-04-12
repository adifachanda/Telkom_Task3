package main

import (
	"fmt"
	"strconv"
)

func generateData(n int) []string {
	baseNumber := "082"
	var customers []string

	for i := 0; i < n; i++ {
		mobileNumber := baseNumber + fmt.Sprintf("%09s", strconv.Itoa(i))
		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscount(input []string) {
	for _, mobileNumber := range input {
		fmt.Println("Sending Promo to", mobileNumber)
	}
	fmt.Println("It's Finished to Send Promo to", len(input), "Customers")
	for _, mobileNumber := range input {
		fmt.Println("Sending Discount to", mobileNumber)
	}
	fmt.Println("It's Finished to Send Discount to", len(input), "Customers")
}

func main() {
	data := generateData(1000)
	sendPromoDiscount(data)
}
