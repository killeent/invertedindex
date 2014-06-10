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
	// recursiveFlag := flag.Bool("recursive", false, "Read files from a directory recursively")
	// flag.Parse()
	// if *recursiveFlag {
	// 	fmt.Println("Working recursively")
	// }

	arguments := os.Args[1:]

	if len(arguments) != 1 {
		usage()
		return
	}

	// Try and get FileInfo
	fileInfo, err := os.Stat(arguments[0])
	if err != nil {
		fmt.Println(err)
		return
	}

	if fileInfo.IsDir() {
		// handle files in directory
		files, err := ioutil.ReadDir(fileInfo.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		for _, fInfo := range files {
			fmt.Printf("Directory %s contains: %s\n", fileInfo.Name(), fInfo.Name())
		}
	} else {
		fmt.Printf("Reading file: %s\n", fileInfo.Name())
		contents, err := ioutil.ReadFile(fileInfo.Name())
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("File %s contains: %s\n", fileInfo.Name(), contents)
	}
}

func usage() {
	fmt.Println("Usage: go run start.go [path]")
}
