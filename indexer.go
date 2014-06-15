package invertedindex

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Indexer struct {
	abort, recursive, verbose bool
	nextDocID                 int
	documents                 map[int]string
	index                     map[string]*list.List
}

func (i *Indexer) BuildIndex(fileInfo os.FileInfo) {
	i.index = make(map[string]*list.List)
	i.documents = make(map[int]string)

	if fileInfo.IsDir() {
		i.readDirectory(fileInfo)
	} else {
		i.readFile(fileInfo)
	}
}

func (i *Indexer) readDirectory(fileInfo os.FileInfo) {
	files, err := ioutil.ReadDir(fileInfo.Name())
	if err != nil && i.abort {
		fmt.Println(err)
		i.cleanup()
		os.Exit(1)
	}
	for _, subFileInfo := range files {
		if subFileInfo.IsDir() {
			if i.recursive {
				i.readDirectory(subFileInfo)
			}
		} else {
			i.readFile(subFileInfo)
		}
	}
}

func (i *Indexer) readFile(fileInfo os.FileInfo) {
	fmt.Printf("Reading file: %s\n", fileInfo.Name())
	contents, err := ioutil.ReadFile(fileInfo.Name())
	if err != nil && i.abort {
		fmt.Println(err)
		i.cleanup()
		os.Exit(1)
	}
	terms := ExtractTerms(contents)
	docID := i.getNextDocID()
	// TODO: need to get absolute path for file in order to index it properly
	i.documents[docID], err = filepath.Abs(fileInfo.Name())
	for _, term := range terms {
		termStr := string(term)
		_, ok := i.index[termStr]
		if !ok {
			i.index[termStr] = list.New()
		}
		i.index[termStr].PushBack(docID)
	}
	fmt.Printf("File %s contains: %s\n", fileInfo.Name(), contents)
}

func (i *Indexer) getNextDocID() int {
	temp := i.nextDocID
	i.nextDocID++
	return temp
}

func (i *Indexer) writeIndexToFile() {

}

// TODO: implement after filewriting;
func (i *Indexer) cleanup() {

}
