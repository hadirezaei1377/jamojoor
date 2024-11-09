package gomock

import (
	"fmt"
	"testing"
)

// mock db in golang

type MockDB struct {
	users map[int]string
}

func (m *MockDB) GetUser(id int) (string, error) {
	if user, ok := m.users[id]; ok {
		return user, nil
	}
	return "", fmt.Errorf("user not found")
}

func TestGetUser(t *testing.T) {
	mockDB := &MockDB{
		users: map[int]string{1: "John Doe"},
	}
	user, err := mockDB.GetUser(1)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if user != "John Doe" {
		t.Errorf("Expected user John Doe, got %v", user)
	}
}
