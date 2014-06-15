/**
Original Author: Trevor Killeen (2014)

Test file for the textProcessor.
*/
package invertedindex

import (
	"bytes"
	"testing"
)

// TODO: refactor unit tests

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

// Tokenization Tests

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

// Normalization Tests

func TestAllLowerNoPunctuation(t *testing.T) {
	tokens := [][]byte{[]byte("abc")}
	expected := [][]byte{[]byte("abc")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestAllUpper(t *testing.T) {
	tokens := [][]byte{[]byte("ABC")}
	expected := [][]byte{[]byte("abc")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemovePeriod(t *testing.T) {
	tokens := [][]byte{[]byte("u.s.a.")}
	expected := [][]byte{[]byte("usa")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemoveComma(t *testing.T) {
	tokens := [][]byte{[]byte("u,s,a,")}
	expected := [][]byte{[]byte("usa")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemoveExclamation(t *testing.T) {
	tokens := [][]byte{[]byte("usa!")}
	expected := [][]byte{[]byte("usa")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemoveQuestion(t *testing.T) {
	tokens := [][]byte{[]byte("?usa")}
	expected := [][]byte{[]byte("usa")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemoveApostrophe(t *testing.T) {
	tokens := [][]byte{[]byte("usa's")}
	expected := [][]byte{[]byte("usas")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemoveDash(t *testing.T) {
	tokens := [][]byte{[]byte("s-i")}
	expected := [][]byte{[]byte("si")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}

func TestRemoveMultiplePunctuationUpperCase(t *testing.T) {
	tokens := [][]byte{[]byte("UsA's-?")}
	expected := [][]byte{[]byte("usas")}
	assertEqualTokenSlices(t, normalize(tokens), expected)
}
