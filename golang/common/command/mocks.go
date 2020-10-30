package command

import "github.com/stretchr/testify/mock"

type mockFileSystemObject struct {
	mock.Mock
}

func (m mockFileSystemObject) Args() []string {
	args := m.Called()
	return args.Get(0).([]string)
}

var mockParams = []string{"cmd", "input-000.txr"}
