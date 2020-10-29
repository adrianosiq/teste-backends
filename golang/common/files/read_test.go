package files

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("Should Read throws if Open throws", func(t *testing.T) {
		mockFileSystem := &MockFileSystemObject{}
		mockFileSystem.On("Open", mockPath).Return(os.Stdin, errors.New("Open throws")).Once()
		fileRead := FileRead{os: mockFileSystem}
		out, err := fileRead.Read(mockPath)
		var a = assert.New(t)
		a.Nil(out)
		a.Error(err)
	})

	t.Run("Should Read return inputs on success", func(t *testing.T) {
		fileRead := FileRead{os: osFS{}}
		out, err := fileRead.Read("mock-input-000.txt")
		var a = assert.New(t)
		a.Nil(err)
		a.NotNil(out)
	})
}
