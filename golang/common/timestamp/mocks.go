package timestamp

import (
	"time"

	"github.com/stretchr/testify/mock"
)

type mockFileSystemObject struct {
	mock.Mock
}

func (m mockFileSystemObject) Parse(value string) (time.Time, error) {
	args := m.Called(value)
	return args.Get(0).(time.Time), args.Error(1)
}

var mockStringTimestamp = "2019-11-11T15:56:04Z"
