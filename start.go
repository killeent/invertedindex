/**
Given a file or directory, reads the contents of that file or the
files in the directory into memory.
*/
package main

import (
	"fmt"
	"os"
)

func main() {
	arguments := os.Args[1:]

	// Attempt to open the file or directory
	file, err := os.Open(os.Args[1])
	if err != nil {
		// err is of type *PathError
		fmt.Println(err)
		return
	}
	// Ensures the file will be closed
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}
	if fileInfo.IsDir() {
		// handle files in directory
	} else {
		fmt.Printf("Reading file: %s\n", fileInfo.Name())
		contents := make([]byte, fileInfo.Size())
		read, err := file.Read(contents)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(contents)
	}
}
