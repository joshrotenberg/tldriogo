package tldrio

import (
	"fmt"
	"testing"
)

func TestCategories(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	categories, err := tldrio.Categories()
	if err != nil {
		t.Error(err)
	} else {
		if len(*categories) == 0 {
			t.Error("no categories found")
		}
	}

}

func ExampleTldrIo_Categories() {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	categories, err := tldrio.Categories()
	if err != nil {
		fmt.Println("couldn't fetch the tl;dr category list")
	} else {
		fmt.Printf("tl;dr has a bunch of categories: %t\n", len(*categories) > 2)
	}
	// Output:
	// tl;dr has a bunch of categories: true

}
