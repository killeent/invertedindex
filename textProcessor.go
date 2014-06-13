/**
Original Author: Trevor Killeen (2014)

This module implements the necessary functions to turn the byte[] of
a text file into a list of terms used in the Inverted Index.
*/

package invertedindex

import (
	"strings"
)

// ExtractTerms takes a byte slice of a text file and parses it into
// a list of terms. To do so it performs tokenization.
func ExtractTerms(file []byte) []string {
	str := string(file)
	tokens := strings.Split(str, "\\s+")
	return tokens
}
