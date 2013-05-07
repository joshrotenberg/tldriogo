package tldrio

import (
	"testing"
)

func TestTldrIo(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()
	t.Logf("%+v", tldrio)
}
