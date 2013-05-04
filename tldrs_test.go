package tldrio

import (
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

func TestSearchNotThere(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	_, err := tldrio.Search("http://not.there.com")
	if err == nil {
		t.Error("expected no results but got one ...")
	}
}
