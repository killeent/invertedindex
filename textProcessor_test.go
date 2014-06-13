package invertedindex

import (
	// "fmt"
	"bytes"
	"testing"
)

// Tokenization Tests

// checks whether two 2-dimensional byte slices are equal
func assertEqualTokenSlices(t *testing.T, a, e [][]byte) {
	if len(a) != len(e) {
		t.Errorf("Expected number of tokens: %d Actual: %d", len(e), len(a))
	}
	for i, token := range a {
		if !bytes.Equal(token, e[i]) {
			t.Error("Tokens not equal")
		}
	}
}

func TestEmptyFile(t *testing.T) {
	bytes := make([]byte, 0)
	assertEqualTokenSlices(t, tokenize(bytes), [][]byte{})
}

func TestWhiteSpace(t *testing.T) {
	bytes := []byte(" ")
	assertEqualTokenSlices(t, tokenize(bytes), [][]byte{})
}

func TestNewLine(t *testing.T) {
	bytes := []byte("\n")
	assertEqualTokenSlices(t, tokenize(bytes), [][]byte{})
}

func TestTab(t *testing.T) {
	bytes := []byte("\t")
	assertEqualTokenSlices(t, tokenize(bytes), [][]byte{})
}

func TestSingleTerm(t *testing.T) {
	bytes := []byte("hi")
	expected := [][]byte{[]byte("hi")}
	assertEqualTokenSlices(t, tokenize(bytes), expected)
}

func TestSingleTermSurroundingWhiteSpace(t *testing.T) {
	bytes := []byte("  hi ")
	expected := [][]byte{[]byte("hi")}
	assertEqualTokenSlices(t, tokenize(bytes), expected)
}

func TestMultipleTerms(t *testing.T) {
	bytes := []byte("hi there")
	expected := [][]byte{[]byte("hi"), []byte("there")}
	assertEqualTokenSlices(t, tokenize(bytes), expected)
}

func TestMultipleTermsWhiteSpace(t *testing.T) {
	bytes := []byte(" hi there  \n")
	expected := [][]byte{[]byte("hi"), []byte("there")}
	assertEqualTokenSlices(t, tokenize(bytes), expected)
}
