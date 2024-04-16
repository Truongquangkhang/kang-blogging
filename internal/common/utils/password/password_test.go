package util_password

import (
	"testing"
)

func Test_HashPassword(t *testing.T) {
	mockPassword := "secret"
	_, err := HashPassword(mockPassword)
	if err != nil {
		t.Fatalf("Get an error: %v", err)
	}
}

func Test_CheckPassword(t *testing.T) {
	mockPassword := "secret"
	hashedPassword, err := HashPassword(mockPassword)
	if err != nil {
		t.Fatalf("Get an error: %v", err)
	}
	if !CheckPasswordHash(mockPassword, hashedPassword) {
		t.Fatalf("Want %s, got %s", mockPassword, hashedPassword)
	}
}
