package recursion

import (
	"fmt"
)

type Item interface {
	IsKey() (int, bool)
	IsBox() (*Box, bool)
}

type Items []Item

type Key int

func (k Key) IsKey() (int, bool) {
	return int(k), true
}

func (k Key) IsBox() (*Box, bool) {
	return nil, false
}

type Box struct {
	Items Items
}

func (b Box) IsKey() (int, bool) {
	return 0, false
}

func (b Box) IsBox() (*Box, bool) {
	return &b, true
}

// LookForKey1 is a non-recursive method for looking through a box with boxes
// for a key
func LookForKey1(mainBox Box) {
	var pile Items = mainBox.Items

	for len(pile) > 0 {
		item := pile[0] // grab item off the top of the pile
		pile = pile[1:] // djust the pile for the item we just took off

		// NOTE, in the book it says `pile.grab_a_box`, it might be assuming
		// that all items in that intial pile are boxes?...
		box, ok := item.IsBox()
		if ok {
			pile = append(pile, box.Items...) // add items to the pile

			continue
		}

		key, ok := item.IsKey()
		if ok {
			fmt.Printf("Found The Key!, %d", key)

			break
		}

	}
}

// LookForKey2 uses recursion to look through boxes to find the key
func LookForKey2(box *Box) {
	for _, item := range box.Items {
		box, ok := item.IsBox()
		if ok {
			LookForKey2(box)

			continue
		}

		key, ok := item.IsKey()
		if ok {
			fmt.Printf("Found The Key!, %d", key)

			break
		}
	}
}

/// infiniteLoop represents an infinite loop through recursion
func infiniteLoop(i int) {
	fmt.Println(i)

	infiniteLoop(i)
}

// Countdown illustrates the base vs recursive case in recursion
func Countdown(i int) {
	fmt.Println(i)

	// base case
	if i <= 0 {
		return
	}

	// rescursive case
	Countdown(i - 1)
}
