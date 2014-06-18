/**
Original Author: Trevor Killeen (2014)

Given a file or directory, reads the contents of that file or the
files in the directory into memory.
*/
package main

import (
	"flag"
	"fmt"
	"github.com/killeent/invertedindex"
	"io/ioutil"
	"os"
)

func main() {
	// Components to be passed to our indexer
	var indexDir string
	var abort, recursive bool

	flag.BoolVar(&abort, "a", false, "If a file or directory cannot be read during indexing"+
		"terminate immediately")
	flag.BoolVar(&recursive, "r", false, "Index the directory contents recursively")
	// flag.BoolVar(&verbose, "v", false, "Log information about the indexing process to to the console")
	// flag.Parse()

	if len(flag.Args()) != 1 {
		usage()
		os.Exit(1)
	}
	indexDir = flag.Args()[0]

	if recursive {
		fmt.Println("Reading the directory recursively")
	}

	// // Try and get FileInfo for the passed directory string
	// fileInfo, err := os.Stat(indexDir)
	// if err != nil {
	// 	// TODO: improve error message
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	indexer := new(Indexer)
	flags := IndexerFlags{Abort: abort, Recursive: recursive}
	indexer.BuildIndex(flags, indexDir)
}

func usage() {
	fmt.Println("Usage: ./start [flags] [directory path]")
}
