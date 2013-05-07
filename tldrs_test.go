package tldrio

import (
	"fmt"
	"testing"
)

func TestLatest(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	latest, err := tldrio.Latest(1, "world-news")
	if err != nil {
		t.Error(err)
	}
	if len((*latest)) != 1 {
		t.Error("didn't get a tl;dr")
	}
}

func ExampleTldrIo_Latest() {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	latest, err := tldrio.Latest(3, "world-news")
	if err != nil {
		fmt.Printf("couldn't get the three latest World News tl;drs: %s\n", err)
	} else {
		fmt.Printf("got three of the latest World News tl;drs: %t\n", len(*latest) == 3)
	}
	// Output:
	// got three of the latest World News tl;drs: true

}

func TestSearch(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	latest, err := tldrio.Latest(1, "world-news")
	if err != nil {
		t.Error(err)
	} else {
		url := (*latest)[0].OriginalUrl
		result, err := tldrio.Search(url)
		if err != nil {
			t.Error(err)
		} else {
			if (*latest)[0].CreatedAt != result.CreatedAt {
				t.Error("tl;drs don't match!")
			}
		}
	}
}

func ExampleTldrIo_Search() {

}


func TestSearchNotThere(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	_, err := tldrio.Search("http://not.there.com")
	if err == nil {
		t.Error("expected no results but got one ...")
	}
}

func TestSearchBatch(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	results, err := tldrio.SearchBatch("http://42floors.com/blog/yc-without-being-in-yc",
		"http://jsonapi.org")
	if err != nil {
		t.Error(err)
	} else {
		if len(*results) == 0 {
			t.Error("batch search didn't return expected results")
		}
	}
}
