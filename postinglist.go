package main

import (
	"container/list"
	"fmt"
)

type posting struct {
	docID     int
	positions *list.List
}

// intersectPostingList returns a list of docIDs that represent the
// intersection of p1 and p2
func intersectPostingList(p1, p2 list.List) *list.List {
	result := list.New()
	e1 := p1.Front()
	e2 := p2.Front()
	for e1 != nil && e2 != nil {
		d1 := e1.Value.(posting)
		d2 := e2.Value.(posting)
		if d1.docID == d2.docID {
			result.PushBack(d1.docID)
			e1 = e1.Next()
			e2 = e2.Next()
		} else if d1.docID < d2.docID {
			e1 = e1.Next()
		} else {
			e2 = e2.Next()
		}
	}
	return result
}

func main() {
	// create two posting lists; we ignore their contents for now
	p1 := list.New()
	p2 := list.New()
	p1.PushBack(posting{docID: 1, positions: nil})
	p1.PushBack(posting{docID: 3, positions: nil})
	p1.PushBack(posting{docID: 5, positions: nil})
	p2.PushBack(posting{docID: 2, positions: nil})
	p2.PushBack(posting{docID: 3, positions: nil})
	p2.PushBack(posting{docID: 4, positions: nil})
	p2.PushBack(posting{docID: 7, positions: nil})

	intersection := intersectPostingList(*p1, *p2)
	for e := intersection.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
