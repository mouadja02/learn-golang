package main

import "fmt"

func main() {
	// initialize variables here
	var smsSendingLimit int = 100
	var costPerSMS float64
	costPerSMS = 0.05
	hasPermission, username := true, "mouad"

	fmt.Printf("Hello %s, your permission status is %v. The cost per sms is %.2f and you have %d sms sending limit\n", username, hasPermission, costPerSMS, smsSendingLimit)
}
