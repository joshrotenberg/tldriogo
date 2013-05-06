package tldrio

import (
	"testing"
)

func TestUser(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	latest, err := tldrio.Latest(1, "world-news")
	if err != nil {
		t.Error(err)
	} else {
		username := (*latest)[0].Creator.Username
		user, err := tldrio.User(username)
		if err != nil {
			t.Error(err)
		} else {
			if username != user.Username {
				t.Error("users don't match")
			}
		}

	}
}

func TestUserTldrs(t *testing.T) {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	latest, err := tldrio.Latest(1, "world-news")
	if err != nil {
		t.Error(err)
	} else {
		username := (*latest)[0].Creator.Username
		tldrs, err := tldrio.UserTldrs(username)
		if err != nil {
			t.Error(err)
		} else {
			t.Logf("%+v", tldrs)
		}

	}
}
