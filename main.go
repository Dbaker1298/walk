package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type config struct {
	// extension to filter out
	ext string

	// min file size
	size int64

	// list files
	list bool

	// delete files
	del bool
}

func main() {
	// Parsing command line flags
	root := flag.String("root", ".", "Root directory to start")
	list := flag.Bool("list", false, "List files only")
	del := flag.Bool("del", false, "Delete files")
	ext := flag.String("ext", "", "File extension to filter out")
	size := flag.Int64("size", 0, "Minimum file size")
	flag.Parse()

	c := config{
		ext:  *ext,
		size: *size,
		list: *list,
		del:  *del,
	}

	if err := run(*root, os.Stdout, c); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func run(root string, out io.Writer, cfg config) error {
	// Walk the directory tree
	return filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		// Check if we have an error
		if err != nil {
			return err
		}

		if filterOut(path, cfg.ext, cfg.size, info) {
			return nil
		}

		// If list was explicitly set, don't do anything else
		if cfg.list {
			return listFile(path, out)
		}

		// Delete files
		if cfg.del {
			return delFile(path)
		}

		// List is the default option if nothing else was set
		return listFile(path, out)
	})
}
