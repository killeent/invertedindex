package invertedindex

import (
	"container/list"
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

type positionalResult struct {
	docID, w1Pos, w2Pos int
}

func abs(n int) int {
	if n < 0 {
		n = -n
	}
	return n
}

// positionalIntersect returns a list of positionalResults, which represent documents
// where both words are present and within k positions of each other. Note that this
// can return duplicate documents when there are multiple places where the two intersect
func positionalIntersect(p1, p2 list.List, k int) *list.List {
	result := list.New()
	e1 := p1.Front()
	e2 := p2.Front()
	for e1 != nil && e2 != nil {
		d1 := e1.Value.(posting)
		d2 := e2.Value.(posting)
		if d1.docID == d2.docID {
			// check positions
			l := list.New()
			pos1 := d1.positions.Front()
			pos2 := d2.positions.Front()
			for pos1 != nil {
				for pos2 != nil {
					if abs(pos1.Value.(int)-pos2.Value.(int)) <= k {
						l.PushBack(pos2.Value.(int))
					} else if pos2.Value.(int) > pos1.Value.(int) {
						break
					}
					pos2 = pos2.Next()
				}
				for l.Len() > 0 &&
					abs(l.Front().Value.(int)-pos1.Value.(int)) > k {
					l.Remove(l.Front())
				}
				for e := l.Front(); e != nil; e = e.Next() {
					result.PushBack(positionalResult{docID: d1.docID, w1Pos: pos1.Value.(int), w2Pos: e.Value.(int)})
				}
				pos1 = pos1.Next()
			}
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

// func main() {
// 	// create two posting lists; we ignore their contents for now
// 	p1 := list.New()
// 	p2 := list.New()
// 	p1.PushBack(posting{docID: 1, positions: nil})
// 	p1.PushBack(posting{docID: 3, positions: nil})
// 	p1.PushBack(posting{docID: 5, positions: nil})
// 	p2.PushBack(posting{docID: 2, positions: nil})
// 	p2.PushBack(posting{docID: 3, positions: nil})
// 	p2.PushBack(posting{docID: 4, positions: nil})
// 	p2.PushBack(posting{docID: 7, positions: nil})

// 	intersection := intersectPostingList(*p1, *p2)
// 	for e := intersection.Front(); e != nil; e = e.Next() {
// 		fmt.Println(e.Value)
// 	}

// 	// now we will concern ourselves with contents
// 	p1.Init()
// 	p2.Init()
// 	post1 := list.New()
// 	post1.PushBack(1)
// 	post1.PushBack(3)
// 	post1.PushBack(6)
// 	post1.PushBack(10)
// 	post2 := list.New()
// 	post2.PushBack(2)
// 	post2.PushBack(5)
// 	post2.PushBack(8)
// 	post2.PushBack(15)
// 	p1.PushBack(posting{docID: 1, positions: post1})
// 	p2.PushBack(posting{docID: 1, positions: post2})

// 	var positionalIntersection *list.List

// 	fmt.Println("First looking for adjacent words (k = 1)")
// 	positionalIntersection = positionalIntersect(*p1, *p2, 1)
// 	for e := positionalIntersection.Front(); e != nil; e = e.Next() {
// 		result := e.Value.(positionalResult)
// 		fmt.Printf("docID: %d, index of first word: %d, index of second word: %d\n",
// 			result.docID, result.w1Pos, result.w2Pos)
// 	}

// 	fmt.Println("Now looking for words within k = 2 positions")
// 	positionalIntersection = positionalIntersect(*p1, *p2, 2)
// 	for e := positionalIntersection.Front(); e != nil; e = e.Next() {
// 		result := e.Value.(positionalResult)
// 		fmt.Printf("docID: %d, index of first word: %d, index of second word: %d\n",
// 			result.docID, result.w1Pos, result.w2Pos)
// 	}

// 	fmt.Println("Now looking for words within k = 5 positions")
// 	positionalIntersection = positionalIntersect(*p1, *p2, 5)
// 	for e := positionalIntersection.Front(); e != nil; e = e.Next() {
// 		result := e.Value.(positionalResult)
// 		fmt.Printf("docID: %d, index of first word: %d, index of second word: %d\n",
// 			result.docID, result.w1Pos, result.w2Pos)
// 	}
// }
