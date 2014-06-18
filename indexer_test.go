package invertedindex

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"sort"
	"testing"
)

// Tests for crawling a directory and properly mapping docIDs
// to paths and handling errors along the way

var emptypath string = "test_files/empty_files"

// tests crawling an empty directory
func TestCrawlEmptyDirectory(t *testing.T) {
	giPath := filepath.Join(emptypath, "empty", ".gitignore")
	gitignore := setupEmptyDirectory(t, giPath)
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "empty"))
	actual := indexer.documents
	expected := map[int]string{}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("improper reading of empty directory")
	}
	teardownEmptyDirectory(t, giPath, gitignore)
}

// git only supports files to be listed in the staging area, not directories, so we
// can't add an empty directory to our git repository. Thus we need to remove our placeholder
// .gitignore file from the empty directory in order to test crawling it; and then put it back
// after we are done

// TODO: improve error messages

// setupEmptyDirectory takes a path to a .gitignore file, reads that file into
// memory and then deletes it. this function assumes that the passed string is the directory
// containing a single .gitignore file; the byte slice containing the file contents are returned
// to the caller so that it can pass them to the teardown method to write them back to file
func setupEmptyDirectory(t *testing.T, path string) []byte {
	gitignore, err := ioutil.ReadFile(path)
	if err != nil {
		t.Error(err)
	}
	os.Remove(path)
	return gitignore
}

// writes the byte slice of the .gitignore file to the passed path
func teardownEmptyDirectory(t *testing.T, path string, contents []byte) {
	err := ioutil.WriteFile(path, contents, 0644)
	if err != nil {
		t.Error(err)
	}
}

// tests crawling a single file
func TestCrawlSingleFile(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "single", "a.txt"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(emptypath, "single", "a.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// tests crawling a flat directory
func TestCrawlFlatDirectory(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "flat"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(emptypath, "flat", "a.txt"),
		1: filepath.Join(emptypath, "flat", "b.txt"),
		2: filepath.Join(emptypath, "flat", "c.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// tests crawling a nested directory without the recursive flag set to false
func TestCrawlNestedDirectoryNonRecursive(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "nested"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(emptypath, "nested", "a.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// tests crawling a nested directory with the recursive flag set to true
func TestCrawlNestedDirectoryRecursive(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{Recursive: true}, filepath.Join(emptypath, "nested"))
	actual := indexer.documents
	expected := map[int]string{0: filepath.Join(emptypath, "nested", "a.txt"),
		1: filepath.Join(emptypath, "nested", "sub1", "b.txt"),
		2: filepath.Join(emptypath, "nested", "sub2", "sub3", "c.txt")}
	assertEqualDocumentMapping(t, actual, expected)
}

// Because git cannot add unreadable files and directories to a repository we temporarily
// make them unreadable for testing purposes and then change their positions back
// after we are done.

func setupUnreadableTestFile() {
	unreadable_file := filepath.Join(emptypath, "unreadable", "unreadable_file.txt")
	os.Chmod(unreadable_file, 0000)
}

func teardownUnreadableTestFile() {
	unreadable_file := filepath.Join(emptypath, "unreadable", "unreadable_file.txt")
	os.Chmod(unreadable_file, 0644)
}

func setupUnreadableTestDir() {
	unreadable_dir := filepath.Join(emptypath, "unreadable", "unreadable_dir")
	os.Chmod(unreadable_dir, 0000)
}

func teardownUnreadableTestDir() {
	unreadable_dir := filepath.Join(emptypath, "unreadable", "unreadable_dir")
	os.Chmod(unreadable_dir, 0755)
}

// tests crawling an unreadable file with the abort flag set to false
func TestCrawlUnreadableFileNotAbort(t *testing.T) {
	setupUnreadableTestFile()
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "unreadable", "unreadable_file.txt"))
	actual := indexer.documents
	expected := map[int]string{}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("improper handling of unreadable file")
	}
	teardownUnreadableTestFile()
}

// tests crawling an unreadable file with the abourt flag set to true
// func TestCrawlUnreadableFileAbort(t *testing.T) {
// 	indexer := setUpIndexer(t, IndexerFlags{Abort: true}, filepath.Join(emptypath, "unreadable", "unreadable_file.txt"))
// }

// tests crawling an unreadable file with the abort flag set to false
func TestCrawlUnreadableDirectoryNotAbort(t *testing.T) {
	setupUnreadableTestDir()
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "unreadable", "unreadable_dir"))
	actual := indexer.documents
	expected := map[int]string{}
	if !reflect.DeepEqual(expected, actual) {
		t.Error("improper handling of unreadable file")
	}
	teardownUnreadableTestDir()
}

// tests crawling an unreadable directory with the abourt flag set to true
// func TestCrawlUnreadableDirectoryAbort(t *testing.T) {
// 	indexer := setUpIndexer(t, IndexerFlags{Abort: true}, filepath.Join(emptypath, "unreadable", "unreadable_dir"))
// }

// Tests for properly building an inverted index mapping terms to docIDs

var indexpath string = "test_files/index_files"

func TestIndexEmptyFile(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "single", "a.txt"))
	actual := indexer.index
	expected := [][]string{}
	assertCorrectIndexMapping(t, actual, expected)
}

func TestIndexEmptyDirectory(t *testing.T) {
	giPath := filepath.Join(emptypath, "empty", ".gitignore")
	gitignore := setupEmptyDirectory(t, giPath)
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(emptypath, "empty"))
	actual := indexer.index
	expected := [][]string{}
	assertCorrectIndexMapping(t, actual, expected)
	teardownEmptyDirectory(t, giPath, gitignore)
}

func TestIndexUniqueWordsInFile(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(indexpath, "unique.txt"))
	actual := indexer.index
	expected := [][]string{[]string{"alpha", "beta", "gamma"}}
	assertCorrectIndexMapping(t, actual, expected)
}

func TestIndexSameWordsInFile(t *testing.T) {
	indexer := setUpIndexer(t, IndexerFlags{}, filepath.Join(indexpath, "duplicate.txt"))
	actual := indexer.index
	expected := [][]string{[]string{"alpha"}}
	assertCorrectIndexMapping(t, actual, expected)
}

func TestIndexSameAndUniqueWords(t *testing.T) {

}

func TestIndexMultipleFiles(t *testing.T) {

}

func setUpIndexer(t *testing.T, flags IndexerFlags, filePath string) *Indexer {
	indexer := new(Indexer)
	// fileInfo := getFileInfo(t, filePath)
	indexer.BuildIndex(flags, filePath)
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

// assertCorrectIndexMapping tests that an index is properly constructed; that is it
// stores a proper mapping of terms to docIDs; We nay not know exactly what docID is
// assigned to a document; so in order to verify that things are equal we rebuild a
// mapping from docID to terms and then check that the individual term lists match
// a slice of expected term slices
func assertCorrectIndexMapping(t *testing.T, actual map[string][]int, expected [][]string) {
	rebuilt := make(map[int][]string)
	for term, docIDs := range actual {
		for docID := range docIDs {
			_, ok := rebuilt[docID]
			if !ok {
				rebuilt[docID] = []string{}
			}
			rebuilt[docID] = append(rebuilt[docID], term)
		}
	}
	if len(rebuilt) != len(expected) {
		t.Errorf("Expected number of documents indexed: %d, actual: %d", len(expected),
			len(rebuilt))
	}
	for _, terms := range rebuilt {
		found := false
		for i := 0; i < len(expected); i++ {
			if reflect.DeepEqual(terms, expected[i]) {
				expected = append(expected[:i], expected[i+1:]...)
				found = true
				break
			}
		}
		if !found {
			t.Error("Some document improperly indexed")
		}
	}
}
