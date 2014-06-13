package invertedindex

import (
"os"
"fmt"
"io/ioutil"
)

func BuildIndex(fileInfo os.FileInfo) {
	if fileInfo.IsDir() {
		readDirectory(fileInfo)
	} else {
		readFile(fileInfo)
	}
}

func readDirectory(fileInfo os.FileInfo) {
	files, err := ioutil.ReadDir(fileInfo.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, subFileInfo := range files {
		if subFileInfo.IsDir() {
			if recursive {
				readDirectory(subFileInfo)
			}
		} else {
			readFile(subFileInfo)
		}
	}
}

func readFile(fileInfo os.FileInfo) {
	fmt.Printf("Reading file: %s\n", fileInfo.Name())
	contents, err := ioutil.ReadFile(fileInfo.Name())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("File %s contains: %s\n", fileInfo.Name(), contents)
}