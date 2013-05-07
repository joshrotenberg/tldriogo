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
	username := "Louis"

	user, err := tldrio.User(username)
	if err != nil {
		fmt.Printf("user %s doesn't exist\n", username)
	} else {
		fmt.Printf("%s was created at %s\n", user.Username, user.CreatedAt)
	}
	// Output:
	// Louis was created at 2012-09-05 08:39:24 +0000 UTC
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

func ExampleTldrIo_UserTldrs() {
	var tldrio *TldrIo
	tldrio = NewTldrIo()

	username := "Louis"

	tldrs, err := tldrio.UserTldrs(username)
	if err != nil {
		fmt.Printf("user %s doesn't exist\n", username)
	} else {
		fmt.Printf("user %s has more than one tldr: %t\n", username, len(*tldrs) > 1)
	}
	// Output:
	// user Louis has more than one tldr: true

}
