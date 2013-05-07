package tldrio

import (
	"fmt"
	"time"
)

type Gravatar struct {
	Url string
}

// User contains information about a TL;DR user.
type User struct {
	Username      string
	TwitterHandle string
	Bio           string
	Gravatar      Gravatar
	CreatedAt     time.Time
	LastActive    time.Time
}

// Struct reprsentation of a single TL;DR
// Note that this data structure is almost identical to Tldr but because of a small inconsistency in the API,
// Creators in the /users/<username>/tldrsCreated call are just a string ID rather than an object.
type UserTldr struct {
	Title                string
	Slug                 string
	Permalink            string
	SummaryBullets       []string
	ReadCount            int
	OriginalUrl          string
	PossibleUrls         []string
	Creator              string
	ImageUrl             string
	CreatedAt            time.Time
	UpdatedAt            time.Time
	Moderated            bool
	DistributionChannels DistributionChannels
	Anonymous            bool
	WordCount            int
	ArticleWordCount     int
	TimeSaved            string
	Language             Language
	Domain               Domain
	Categories           []Category
	Editors              []string
	ThankedBy            []string
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
func (t *TldrIo) UserTldrs(user string) (*[]UserTldr, error) {
	var tldrs []UserTldr
	err := callApi(t, "GET", fmt.Sprintf("users/%s/tldrsCreated", user), "", nil, &tldrs)
	if err != nil {
		return nil, err
	}
	return &tldrs, nil
}
