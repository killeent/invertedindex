package invertedindex

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
)

type Indexer struct {
	abort, recursive bool
	index            map[string]*list.List
}

func (i *Indexer) BuildIndex(fileInfo os.FileInfo) {
	i.index = make(map[string]*list.List)

	if fileInfo.IsDir() {
		readDirectory(fileInfo)
	} else {
		readFile(fileInfo)
	}
}

func (i *Indexer) readDirectory(fileInfo os.FileInfo) {
	files, err := ioutil.ReadDir(fileInfo.Name())
	if err != nil && i.abort {
		fmt.Println(err)
		os.Exit(1)
	}
	for _, subFileInfo := range files {
		if subFileInfo.IsDir() {
			if i.recursive {
				readDirectory(subFileInfo)
			}
		} else {
			readFile(subFileInfo)
		}
	}
}

func (i *Indexer) readFile(fileInfo os.FileInfo) {
	fmt.Printf("Reading file: %s\n", fileInfo.Name())
	contents, err := ioutil.ReadFile(fileInfo.Name())
	if err != nil && i.abort {
		fmt.Println(err)
		os.Exit(1)
	}
	terms := ExtractTerms(contents)
	for _, term := range terms {
	}
	fmt.Printf("File %s contains: %s\n", fileInfo.Name(), contents)
}

func (i *Indexer) getNextDocID() int {

}

func (i *Indexer) writeIndexToFile() {

}

// TODO: implement after filewriting;
func (i *Indexer) cleanup() {

}
