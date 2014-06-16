package invertedindex

import (
	"fmt"
	// "os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

// Tests for crawling a directory and properly mapping docIDs
// to paths and handling errors along the way

var crawlpath string = "test_files/empty_files"

// tests crawling an empty directory
func TestCrawlEmptyDirectory(t *testing.T) {
	indexer := setUpIndexer(t, filepath.Join(crawlpath, "empty"))
	actual := indexer.documents
	expected := map[int]string{}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("docID to document mapping invalid")
	}
}

// tests crawling a single file
func TestCrawlSingleFile(t *testing.T) {
	indexer := setUpIndexer(t, filepath.Join(crawlpath, "single", "a.txt"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(crawlpath, "single", "a.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// tests crawling a flat directory
func TestCrawlFlatDirectory(t *testing.T) {
	indexer := setUpIndexer(t, filepath.Join(crawlpath, "flat"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(crawlpath, "flat", "a.txt"),
		1: filepath.Join(crawlpath, "flat", "b.txt"),
		2: filepath.Join(crawlpath, "flat", "c.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// tests crawling a nested directory without the recursive flag set
func TestCrawlNestedDirectoryNonRecursive(t *testing.T) {
	indexer := setUpIndexer(t, filepath.Join(crawlpath, "nested"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(crawlpath, "nested", "a.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// tests crawling a nested directory with the recursive flag set
func TestCrawlNestedDirectoryRecursive(t *testing.T) {
	// indexer := setUpIndexer(t, filepath.Join(crawlpath, "nested"))
}

func setUpIndexer(t *testing.T, filePath string) *Indexer {
	indexer := new(Indexer)
	// fileInfo := getFileInfo(t, filePath)
	indexer.BuildIndex(filePath)
	return indexer
}

// getFileInfo tries to get and return fileInfo struct for the passed file
// path. If the file does not exists, this function raises an error on the
// testing framework
// func getFileInfo(t *testing.T, path string) os.FileInfo {
// 	path, err := filepath.Abs(path)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	fileInfo, err := os.Stat(path)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	return fileInfo
// }

// assertEqualDocumentMapping tests that two documents maps (docID -> file path)
// are equivalent, that is they store the same docIDs and same file paths. They
// specific mapping of docID to path is not important because it is not guaranteed
// to be the same every time
func assertEqualDocumentMapping(t *testing.T, actual, expected map[int]string) {
	if len(actual) != len(expected) {
		t.Errorf("Expected number of documents indexed: %d, actual: %d", len(expected),
			len(actual))
	}
	actualKeys := make([]int, len(actual))
	actualPaths := make([]string, len(actual))
	expectedKeys := make([]int, len(expected))
	expectedPaths := make([]string, len(expected))
	i := 0
	for k, v := range actual {
		actualKeys[i] = k
		actualPaths[i] = v
		fmt.Println(v)
		i++
	}
	i = 0
	for k, v := range expected {
		expectedKeys[i] = k
		expectedPaths[i] = v
		fmt.Println(v)
		i++
	}
	sort.Ints(actualKeys)
	sort.Ints(expectedKeys)
	if !reflect.DeepEqual(actualKeys, expectedKeys) {
		t.Error("Invalid docIDs indexed")
	}
	sort.Strings(actualPaths)
	sort.Strings(expectedPaths)
	if !reflect.DeepEqual(actualPaths, expectedPaths) {
		t.Error("Invalid paths indexed")
	}
}
