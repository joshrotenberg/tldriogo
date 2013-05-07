package tldrio

import (
	"fmt"
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
			// TODO: validate all fields
			if username != user.Username {
				t.Error("users don't match")
			}
		}

	}
}

func ExampleTldrIo_User() {
	var tldrio *TldrIo
	tldrio = NewTldrIo()
	username := "joshrotenberg"

	user, err := tldrio.User(username)
	if err != nil {
		fmt.Printf("user %s doesn't exist\n", username)
	} else {
		fmt.Printf("%s was created at %s\n", user.Username, user.CreatedAt)
	}
	// Output:
	// JoshRotenberg was created at 2013-04-12 17:16:10.363 +0000 UTC
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
			// TODO: validate all fields
			if len(*tldrs) == 0 {
				t.Errorf("expected >0 tl;drs by user %s", username)
			}
		}

	}
}
