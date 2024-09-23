package database

import "github.com/stretchr/testify/mock"

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) DbPingMethod() error {
	args := m.Called()
	return args.Error(0)
}
func (m *MockUserRepository) RegisterUserMethod(username, password string) (int, error) {
	args := m.Called(username, password)
	return args.Int(0), args.Error(1)
}
