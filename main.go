package main

import (
	"fmt"
	"os"

	flags "github.com/jessevdk/go-flags"

	"github.com/ibbbpbbbp/H264analyzer/parser"
	"github.com/ibbbpbbbp/gobits"
)

func main() {
	var opts struct {
		PositionalArgs struct {
			File string `description:"File of H.264 Annex B byte stream format."`
		} `positional-args:"yes" required:"yes"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		os.Exit(1)
	}

	file, err := os.Open(opts.PositionalArgs.File)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	defer file.Close()

	ba := gobits.NewIOByteAccessor(file)
	parser.ParseH264Stream(ba)
}
