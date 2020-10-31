package timestamp

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringToTimestamp(t *testing.T) {
	t.Run("Should ParseTimestamp throws if Parse throws", func(t *testing.T) {
		mockFileSystem := &mockFileSystemObject{}
		mockFileSystem.On("Parse", mockStringTimestamp).Return(time.Time{}, errors.New("Parse throws")).Once()
		stringToTimestamp := StringToTimestamp{os: mockFileSystem}
		out, err := stringToTimestamp.ParseTimestamp(mockStringTimestamp)
		var a = assert.New(t)
		a.Equal(time.Time{}, out)
		a.Error(err)
	})
}
