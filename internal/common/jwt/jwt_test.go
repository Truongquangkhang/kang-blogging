package jwt

import (
	"fmt"
	"testing"
)

func Test_genAccessToken(t *testing.T) {
	mockUserId := "khang"
	token, err := CreateAccessToken(mockUserId, "key", 2)
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
	fmt.Println(token)
}

func Test_VerifyToken(t *testing.T) {
	mockUserId := "khang"
	token, err := CreateAccessToken(mockUserId, "key", 2)
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
	err = VerifyToken(token, "key")
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
}
