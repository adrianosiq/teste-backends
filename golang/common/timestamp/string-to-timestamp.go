package timestamp

import (
	"time"
)

const (
	format = time.RFC3339
)

type fileSystem interface {
	Parse(value string) (time.Time, error)
}

type osFS struct{}

func (osFS) Parse(value string) (time.Time, error) {
	return time.Parse(format, value)
}

//StringToTimestamp is responsible for load services
type StringToTimestamp struct {
	os fileSystem
}

//NewStringToTimestamp create a new string to timestamp
func NewStringToTimestamp() StringToTimestamp {
	return StringToTimestamp{
		os: osFS{},
	}
}

//ParseTimestamp is responsible for convert string to timestamp
func (stringToTimestamp StringToTimestamp) ParseTimestamp(value string) (time.Time, error) {
	timestamp, err := stringToTimestamp.os.Parse(value)
	if err != nil {
		return timestamp, err
	}
	return timestamp, nil
}
