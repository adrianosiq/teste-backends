package files

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("Should Exec throws if Open throws", func(t *testing.T) {
		mockFileSystem := &mockFileSystemObject{}
		mockFileSystem.On("Open", mockPath).Return(os.Stdin, errors.New("Open throws")).Once()
		read := Read{os: mockFileSystem}
		out, err := read.Exec(mockPath)
		var a = assert.New(t)
		a.Nil(out)
		a.Error(err)
	})

	t.Run("Should Exec return inputs on success", func(t *testing.T) {
		read := Read{os: osFS{}}
		out, err := read.Exec("mock-input-000.txt")
		var a = assert.New(t)
		a.Nil(err)
		a.NotNil(out)
	})
}
