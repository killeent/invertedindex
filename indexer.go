package invertedindex

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Indexer struct {
	flags     IndexerFlags
	nextDocID int
	documents map[int]string
	index     map[string][]int
}

type IndexerFlags struct {
	Abort     bool
	Recursive bool
	Verbose   bool
}

// eventual use for testing aborts in code
// type AbortHandler interface {
// 	Abort()
// }

func (i *Indexer) BuildIndex(flags IndexerFlags, path string) {
	fmt.Printf("building index on directory: %s\n", path)
	fileInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	i.flags = flags
	i.index = make(map[string][]int)
	i.documents = make(map[int]string)

	if fileInfo.IsDir() {
		i.readDirectory(fileInfo, path)
	} else {
		i.readFile(fileInfo, filepath.Dir(path))
	}
}

func (i *Indexer) readDirectory(fileInfo os.FileInfo, path string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		if i.flags.Abort {
			fmt.Println(err)
			i.cleanup()
			os.Exit(1)
		} else {
			return
		}
	}
	for _, subFileInfo := range files {
		if subFileInfo.IsDir() {
			if i.flags.Recursive {
				i.readDirectory(subFileInfo, filepath.Join(path, subFileInfo.Name()))
			}
		} else {
			i.readFile(subFileInfo, path)
		}
	}
}

func (i *Indexer) readFile(fileInfo os.FileInfo, dir string) {
	fmt.Printf("Reading file: %s\n", fileInfo.Name())
	contents, err := ioutil.ReadFile(filepath.Join(dir, fileInfo.Name()))
	if err != nil {
		if i.flags.Abort {
			fmt.Println(err)
			i.cleanup()
			os.Exit(1)
		} else {
			return
		}
	}
	terms := ExtractTerms(contents)
	docID := i.getNextDocID()
	i.documents[docID] = filepath.Join(dir, fileInfo.Name())
	for _, term := range terms {
		termStr := string(term)
		_, ok := i.index[termStr]
		if !ok {
			i.index[termStr] = []int{}
		}
		found := false
		for _, id := range i.index[termStr] {
			if id == docID {
				found = true
				break
			}
		}
		if !found {
			// fmt.Printf("adding term: %s id: %d pair to index\n", termStr, docID)
			i.index[termStr] = append(i.index[termStr], docID)
		}
	}
	// fmt.Printf("File %s contains: %s\n", fileInfo.Name(), contents)
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
