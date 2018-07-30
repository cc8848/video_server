package model

import (
	"testing"
)

func TestUserWorkFlow(t *testing.T) {
	clearTable()
	t.Run("AddUser", testAddUserCredential)
	t.Run("GetUserInfo", testGetUserCredential)
	t.Run("DeleteUser", testDeleteUser)
	t.Run("ReGetUserInfo", testReGetUserCredential)
}

func testAddUserCredential(t *testing.T) {
	err := AddUserCredential("Zereker", "jiayou")
	if err != nil {
		t.Errorf("Error of AddUser: %v", err)
	}
}

func testGetUserCredential(t *testing.T) {
	password, err := GetUserCredential("Zereker")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
	if password != "jiayou" {
		t.Errorf("Expect get %s, but get %s", "jiayou", password)
	}
}

func testDeleteUser(t *testing.T) {
	err := DeleteUser("Zereker", "jiayou")
	if err != nil {
		t.Errorf("Error of DeleteUser: %v", err)
	}
}

func testReGetUserCredential(t *testing.T) {
	password, err := GetUserCredential("Zereker")
	if err != nil {
		t.Errorf("Error of GetUser: %v", err)
	}
	if password != "" {
		t.Errorf("Expect get %s, but get %s", "", password)
	}
}
