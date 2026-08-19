package main

import "fmt"

type CreatorSignup struct{ Email, CreatorName, Token string }

func verificationMessage(signup CreatorSignup) string {
	return fmt.Sprintf("Hello %s, verify your creator account with token %s.", signup.CreatorName, signup.Token)
}

func sendVerificationLink(signup CreatorSignup) (EmailResult, error) {
	if signup.Email == "" || signup.CreatorName == "" || signup.Token == "" {
		return EmailResult{}, fmt.Errorf("email, creator name, and token are required")
	}
	return sendEmail(map[string]string{"to": signup.Email, "subject": "Verify your creator account", "body": verificationMessage(signup)}, "creator-verification-"+signup.Email+"-"+signup.Token)
}
