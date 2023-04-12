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

func sendPromoDiscount(input1 []string, input2 []string) {
	for _, mobileNumber := range input1 {
		fmt.Println("Sending Promo to", mobileNumber)
	}
	fmt.Println("It's Finished to Send Promo to", len(input1), "Customers")
	for _, mobileNumber := range input2 {
		fmt.Println("Sending Discount to", mobileNumber)
	}
	fmt.Println("It's Finished to Send Discount to", len(input2), "Customers")
}

func main() {
	dataPromo := generateData(100000000)
	dataDiscount := generateData(1000)
	sendPromoDiscount(dataPromo, dataDiscount)
}
