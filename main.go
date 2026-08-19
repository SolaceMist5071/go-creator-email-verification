package main

import (
	"fmt"
	"os"
)

func main() {
	to := os.Getenv("CREATOR_EMAIL")
	if to == "" {
		panic("CREATOR_EMAIL is required")
	}
	result, err := sendVerificationLink(CreatorSignup{Email: to, CreatorName: "New creator", Token: "signup-token"})
	if err != nil {
		panic(err)
	}
	fmt.Println("verification email sent:", result.MessageID)
}
