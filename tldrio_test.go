package tldrio

import (
	"testing"
)

func TestTldr(t *testing.T) {
	var tldr *Tldr
	tldr = NewTldr()
	t.Logf("%+v", tldr)
}

func TestCategories(t *testing.T) {
	var tldr *Tldr
	tldr = NewTldr()

	tldr.Categories()

}
