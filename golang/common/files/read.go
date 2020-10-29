package files

import (
	"bufio"
	"io"
	"os"
)

var fs fileSystem = osFS{}

type file interface {
	io.Closer
	io.Reader
	io.ReaderAt
	io.Seeker
	Stat() (os.FileInfo, error)
}

type fileSystem interface {
	Open(name string) (file, error)
}

type osFS struct{}

func (osFS) Open(name string) (file, error) {
	return os.Open(name)
}

//FileRead is responsible for load services
type FileRead struct {
	os fileSystem
}

//Read is responsible for read file and return the bytes
func (fileRead FileRead) Read(path string) ([]string, error) {
	handle, err := fileRead.os.Open(path)
	if err != nil {
		return nil, err
	}
	defer handle.Close()
	scanner := bufio.NewScanner(handle)
	scanner.Split(bufio.ScanLines)
	var inputs []string
	for scanner.Scan() {
		inputs = append(inputs, scanner.Text())
	}
	return inputs, nil
}
