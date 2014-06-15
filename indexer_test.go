package invertedindex

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

// Tests for crawling a directory and properly mapping docIDs
// to paths and handling errors along the way

var crawlpath string = "test_files/empty_files"

func getFileInfo(t *testing.T, path string) os.FileInfo {
	path, err := filepath.Abs(path)
	if err != nil {
		t.Error(err)
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		t.Error(err)
	}
	return fileInfo
}

func TestCrawlEmptyDirectory(t *testing.T) {
	indexer := new(Indexer)
	fileInfo := getFileInfo(t, filepath.Join(crawlpath, "empty"))
	indexer.BuildIndex(fileInfo)
	actual := indexer.documents
	expected := make(map[int]string)
	if !reflect.DeepEqual(expected, actual) {
		t.Error("docID to document mapping invalid")
	}
}

// func TestCrawlSingleFile(t *testing.T) {
// 	indexer := new(Indexer)
// }
