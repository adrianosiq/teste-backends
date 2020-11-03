package command

import (
	"errors"
	"os"
)

var fs fileSystem = osFS{}

type fileSystem interface {
	Args() []string
}

type osFS struct{}

func (osFS) Args() []string {
	return os.Args
}

//Arguments is responsible for load services
type Arguments struct {
	os     fileSystem
	params []string
}

//NewArguments create a new arguments
func NewArguments() Arguments {
	return Arguments{
		os: osFS{},
	}
}

//Get is responsible for get arguments line-command
func (arguments *Arguments) Get() ([]string, error) {
	arguments.params = arguments.os.Args()
	if len(arguments.params) < 2 {
		return nil, errors.New("no parameters found")
	}
	return arguments.params, nil
}
