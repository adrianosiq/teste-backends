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

//Read is responsible for load services
type Read struct {
	os fileSystem
}

//NewRead create a new read
func NewRead() Read {
	return Read{
		os: osFS{},
	}
}

//Exec is responsible for read file and return the bytes
func (read Read) Exec(path string) ([]string, error) {
	handle, err := read.os.Open(path)
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
