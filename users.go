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
	Username      string    `json:"username"`
	TwitterHandle string    `json:"twitterHandle"`
	Bio           string    `json:"bio"`
	Gravatar      Gravatar  `json:"gravatar"`
	CreatedAt     time.Time `json:"createdAt"`
	LastActive    time.Time `json:"lastActive"`
}

// Struct reprsentation of a single TL;DR
// Note that this data structure is almost identical to Tldr but because of a small inconsistency in the API,
// Creators in the /users/<username>/tldrsCreated call are just a string ID rather than an object.
type UserTldr struct {
	Title                string               `json:"title"`
	Slug                 string               `json:"slug"`
	Permalink            string               `json:"permalink"`
	SummaryBullets       []string             `json:"summaryBullets"`
	ReadCount            int                  `json:"readCount"`
	OriginalUrl          string               `json:"originalUrl"`
	PossibleUrls         []string             `json:"possibleUrls"`
	Creator              string               `json:"creator"`
	ImageUrl             string               `json:"imageUrl"`
	CreatedAt            time.Time            `json:"createdAt"`
	UpdatedAt            time.Time            `json:"updatedAt"`
	Moderated            bool                 `json:"moderated"`
	DistributionChannels DistributionChannels `json:"distributionChannels"`
	Anonymous            bool                 `json:"anonymous"`
	WordCount            int                  `json:"wordCount"`
	ArticleWordCount     int                  `json:"articleWordCount"`
	TimeSaved            string               `json:"timeSaved"`
	Language             Language             `json:"language"`
	Domain               Domain               `json:"domain"`
	Categories           []Category           `json:"categories"`
	Editors              []string             `json:"editors"`
	ThankedBy            []string             `json:"thankedBy"`
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
