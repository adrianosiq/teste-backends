package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/adrianosiq/teste-backends/golang/cmd/main/factories"
	"github.com/adrianosiq/teste-backends/golang/common/files"

	"github.com/adrianosiq/teste-backends/golang/common/command"
)

func main() {
	arguments := command.NewArguments()
	params, err := arguments.Get()
	if err != nil {
		log.Fatalf(err.Error())
	}
	read := files.NewRead()
	inputs, err := read.Exec(params[1])
	if err != nil {
		log.Fatalf(err.Error())
	}
	out, err := factories.MessagesFactory(inputs)
	if err != nil {
		log.Fatalf(err.Error())
	}
	prettyJSON, err := json.MarshalIndent(out, "", "  ")
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Printf("%s\n", string(prettyJSON))
}
