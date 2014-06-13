package invertedindex

import (
	// "fmt"
	"testing"
)

func TestEmptyFile(t *testing.T) {
	bytes := make([]byte, 0)
	tokens := ExtractTerms(bytes)
	if len(tokens) != 0 {
		t.Error("Empty byte slice should return no tokens")
	}
}

// Tokenization Tests

func TestWhiteSpace(t *testing.T) {
	bytes := []byte(" ")
	tokens := ExtractTerms(bytes)
	if len(tokens) != 0 {
		t.Error("Whitespace only byte slice should return no tokens")
	}
}

func TestSingleTerm(t *testing.T) {
	bytes := []byte("hi")
	tokens := ExtractTerms(bytes)
	if len(tokens) != 1 {
		t.Error("Single term byte slice should return a single token")
	}
}
