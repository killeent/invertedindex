/**
Given a file or directory, reads the contents of that file or the
files in the directory into memory.
*/
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var recursive bool
	var indexDir string
	flag.BoolVar(&recursive, "r", false, "Index the directory contents recursively")
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
		fmt.Println(err)
		return
	}

	if fileInfo.IsDir() {
		// handle files in directory
		files, err := ioutil.ReadDir(fileInfo.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		for _, fInfo := range files {
			fmt.Printf("Directory %s contains: %s\n", fileInfo.Name(), fInfo.Name())
		}
	} else {
		fmt.Printf("Reading file: %s\n", fileInfo.Name())
		contents, err := ioutil.ReadFile(fileInfo.Name())
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Printf("File %s contains: %s\n", fileInfo.Name(), contents)
	}
}

func usage() {
	fmt.Println("Usage: ./start [flags] [directory path]")
}
