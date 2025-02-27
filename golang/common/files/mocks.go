package files

import "github.com/stretchr/testify/mock"

type mockFileSystemObject struct {
	mock.Mock
}

func (m mockFileSystemObject) Open(name string) (file, error) {
	args := m.Called(name)
	return args.Get(0).(file), args.Error(1)
}

var mockPath = "input-000-test.txt"
