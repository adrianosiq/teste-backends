package command

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRead(t *testing.T) {
	t.Run("Should Get throws if no parameters found", func(t *testing.T) {
		mockFileSystem := &mockFileSystemObject{}
		mockFileSystem.On("Args").Return([]string{}).Once()
		arguments := Arguments{os: mockFileSystem}
		out, err := arguments.Get()
		var a = assert.New(t)
		a.Nil(out)
		a.Error(err)
	})

	t.Run("Should Get return parameters on success", func(t *testing.T) {
		arguments := NewArguments()
		out, err := arguments.Get()
		var a = assert.New(t)
		a.Nil(err)
		a.NotNil(out)
	})
}
