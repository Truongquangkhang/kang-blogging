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
	expectedID := "90debe20-1284-4643-a383-ada0e715dafe"
	expectedRole := "user"
	mockSecretKey := "abc"
	token, _ := CreateAccessToken(expectedID, expectedRole, mockSecretKey, 2)

	actualID, actualRole, err := GetIDAndRoleFromJwtToken(token, mockSecretKey)
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
