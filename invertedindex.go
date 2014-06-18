package invertedindex

type InvertedIndex struct {
	// postingDict maps termIDs to the pair (frequency, offset). Frequency is
	// the number of times a word occurs in the index's document collection.
	// Offset is the position of the posting list representation of the
	// associated docIDs for that term in the index file.
	postingDict map[int]Pair

	// termDict maps terms to their termIDs.
	termDict map[string]int

	// docDict maps document paths to docIDs.
	docDict map[string]int
}

// AddDocuments takes a path to a file and adds its contents to the index.
// It assigns a docID to the file, reads in its contents, extracts terms from
// those contents and indexes them.
func (i *InvertedIndex) AddDocument(path string) {

}

// assignDocID generates a unique docID for the associated document indicated
// by path and adds it to the path->docID dictionary.
func assignDocID(path string) {

}

// assignTermID generates a unique termID for the associated term and adds it
// to the term->termID dictionary
func assignTermID(term string) {

}

// WriteToFile writes the contents of the InvertedIndex to a new file
// specified by path. Returns an error if the index cannot be written.
func (i *InvertedIndex) WriteToFile(path string) {

}

// RebuildFromFile generates an InvertedIndex from the index file specified
// by path. Returns an error if an index cannot be constructed from the
// associated file.
func (i *InvertedIndex) RebuildFromFile(path string) error {

}

// Merge combines two index files specified by i1 and i2 and writes them
// to a combined index file in location dest. Returns an error if the
// merge cannot be performed.
func Merge(i1, i2, dest string) error {

}
