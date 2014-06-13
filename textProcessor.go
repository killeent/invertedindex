/**
Original Author: Trevor Killeen (2014)

This module implements the necessary functions to turn the byte[] of
a text file into a list of terms used in the Inverted Index.
*/

package invertedindex

import (
	"bytes"
	// "fmt"
	// "strings"
)

// ExtractTerms takes a byte slice of a text file and parses it into
// a slice of term slices. To do so it performs tokenization.
func ExtractTerms(file []byte) [][]byte {
	return tokenize(file)
}

// tokenize tokenizes a byte slice of text by whitespace and returns
// a slice of slice tokens of terms in the text. If the byte slice is
// empty or is only whitespace, returns an empty slice
func tokenize(file []byte) [][]byte {
	return bytes.Fields(file)
}
