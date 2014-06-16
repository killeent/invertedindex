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

func (i *Indexer) BuildIndex(path string) {
	fmt.Printf("building index on directory: %s\n", path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i.index = make(map[string]*list.List)
	i.documents = make(map[int]string)

	if fileInfo.IsDir() {
		i.readDirectory(fileInfo, path)
	} else {
		i.readFile(fileInfo, filepath.Dir(path))
	}
}

func (i *Indexer) readDirectory(fileInfo os.FileInfo, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil && i.abort {
		fmt.Println(err)
		i.cleanup()
		os.Exit(1)
	}
	for _, subFileInfo := range files {
		if subFileInfo.IsDir() {
			if i.recursive {
				i.readDirectory(subFileInfo, filepath.Join(path, subFileInfo.Name()))
			}
		} else {
			i.readFile(subFileInfo, path)
		}
	}
}

func (i *Indexer) readFile(fileInfo os.FileInfo, dir string) {
	fmt.Printf("Reading file: %s\n", fileInfo.Name())
	contents, err := ioutil.ReadFile(fileInfo.Name())
	if err != nil && i.abort {
		fmt.Println(err)
		i.cleanup()
		os.Exit(1)
	}
	terms := ExtractTerms(contents)
	docID := i.getNextDocID()
	i.documents[docID] = filepath.Join(dir, fileInfo.Name())
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
