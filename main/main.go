/**
Original Author: Trevor Killeen (2014)

Given a file or directory, reads the contents of that file or the
files in the directory into memory.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"github.com/killeent/invertedindex"
)

var recursive bool

func main() {
	var indexDir string
	flag.BoolVar(&recursive, "r", false, "Index the directory contents recursively")
	// flag.BoolVar(&abort, "a", false, "If a file or directory cannot be read during indexing"+
	//	"terminate immediately")
	flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
		os.Exit(1)
	}
	indexDir = flag.Args()[0]

	if recursive {
		fmt.Println("Reading the directory recursively")
	}

	// Try and get FileInfo for the passed directory string
	fileInfo, err := os.Stat(indexDir)
	if err != nil {
		// TODO: improve error message
		fmt.Println(err)
		os.Exit(1)
	}

	invertedindex.BuildIndex(fileInfo)
}

func usage() {
	fmt.Println("Usage: ./start [flags] [directory path]")
}
