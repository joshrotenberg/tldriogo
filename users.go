package tldrio

import (
	"fmt"
	"time"
)

type Gravatar struct {
	Url string `json:"url"`
}

// User contains information about a TL;DR user.
type User struct {
	Username string `json:"username"`
	TwitterHandle string `json:"twitterHandle"`
	Bio string `json:"bio"`
	Gravatar Gravatar `json:"gravatar"`
	CreatedAt time.Time `json:"createdAt"`
	LastActive time.Time `json:"lastActive"`
}

// User gets public data about the given user.
func (t *TldrIo) User(user string) (*User, error) {
	var u User
	err := callApi(t, "GET", fmt.Sprintf("users/%s", user), "", nil, &u)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

// UserTldrs gets the TL;DRs created by the given user.
func (t *TldrIo) UserTldrs(user string) (*[]Tldr, error) {
	var tldrs []Tldr
	err := callApi(t, "GET", fmt.Sprintf("users/%s/tldrsCreated", user), "", nil, &tldrs)
	if err != nil {
		return nil, err
	}
	return &tldrs, nil
}
