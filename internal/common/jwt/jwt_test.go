package jwt

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_genAccessToken(t *testing.T) {
	mockUserId := "khang"
	token, err := CreateAccessToken(mockUserId, "user", "key", 2)
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
	fmt.Println(token)
}

func Test_VerifyToken(t *testing.T) {
	mockUserId := "khang"
	token, err := CreateAccessToken(mockUserId, "user", "key", 2)
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
	err = VerifyToken(token, "key")
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
}

func Test_GetIDAndRoleFromJwtToke(t *testing.T) {
	mockJwtToken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTQ2NDU5MzgsInJvbGUiOiJ1c2VyIiwidXNlcklkIjoiOTBkZWJlMjAtMTI4NC00NjQzLWEzODMtYWRhMGU3MTVkYWZlIn0.1KMLPNS2ylTTOC7wCyjRBEARrYphDi341zp8f89cuAA"
	expectedID := "90debe20-1284-4643-a383-ada0e715dafe"
	expectedRole := "user"

	actualID, actualRole, err := GetIDAndRoleFromJwtToken(mockJwtToken, "abc")
	if err != nil {
		t.Fatalf("Get an error %v", err)
	}
	if !reflect.DeepEqual(actualID, expectedID) {
		t.Fatalf("Wanted %v, got %v", expectedID, actualID)
	}
	if !reflect.DeepEqual(actualRole, expectedRole) {
		t.Fatalf("Wanted %v, got %v", expectedRole, actualRole)
	}

}
