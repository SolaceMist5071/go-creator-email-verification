package main

import "testing"

func TestVerificationMessageUsesCreatorIdentityAndToken(t *testing.T) {
	got := verificationMessage(CreatorSignup{Email: "creator@example.com", CreatorName: "Mina", Token: "abc123"})
	want := "Hello Mina, verify your creator account with token abc123."
	if got != want {
		t.Fatalf("message = %q, want %q", got, want)
	}
}
