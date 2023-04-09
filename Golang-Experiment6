package main

import (
	"fmt"
	"time"
)

const address string = "DKI Jakarta"

func findAddress(addresses [100000]string) {
	tx := time.Now()
	fmt.Println("The Default Address is", addresses[0])
	ty := time.Since(tx)
	fmt.Println("The performance is", ty.Milliseconds(), "ms")
}

func main() {
	var addresses [100000]string
	for i := 0; i < len(addresses); i++ {
		addresses[i] = address
	}
	findAddress(addresses)
}
