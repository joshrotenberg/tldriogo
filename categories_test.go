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
		t.Logf("%+v", categories)
	}

}
