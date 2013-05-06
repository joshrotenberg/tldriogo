package tldrio

import (
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
