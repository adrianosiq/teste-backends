package files

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFileSystemObject struct {
	mock.Mock
}

func (m MockFileSystemObject) Open(name string) (file, error) {
	args := m.Called(name)
	return args.Get(0).(file), args.Error(1)
}

func TestRead(t *testing.T) {
	t.Run("Should Read throws if Open throws", func(t *testing.T) {
		mockPath := "input-000-test.txt"
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
