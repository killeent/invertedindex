package invertedindex

import (
	"io/ioutil"
	"testing"
)

func TestEmptyFile(t *testing.T) {
	contents, _ := ioutil.ReadFile("test_files/empty.txt")
	tokens := ExtractTerms(contents)
	if len(tokens) != 0 {
		t.Errorf("Empty file tokenized improperly. Expected: %d tokens; Actual: %d tokens", 0, len(tokens))
	}
}
