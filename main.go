package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/gogo/protobuf/plugin/size"
)

type config struct {
	// extension to filter out
	ext string

	// min file size
	size int64

	// list files
	list bool
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}