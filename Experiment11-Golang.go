package main

import "fmt"

func generateData() []string {
	number := "0821234567"
	customers := []string{}
	var mobileNumber string

	for i := 0; i < 100; i++ {
		if i < 10 {
			mobileNumber = number + "0" + fmt.Sprintf("%d", i)
		} else {
			mobileNumber = number + fmt.Sprintf("%d", i)
		}
		customers = append(customers, mobileNumber)
	}
	return customers
}

func sendPromoDiscount(array []string) {
	for i := 0; i < len(array); i++ {
		fmt.Println("Sending promo to", array[i])
	}
	for i := 0; i < 10; i++ {
		fmt.Println("Sending Discount to Chosen Customer", i+1)
	}
}

func main() {
	customers := generateData()
	sendPromoDiscount(customers)
}
