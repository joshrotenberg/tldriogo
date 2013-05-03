package tldrio

import (
	"fmt"
	"net/url"
	"time"
)

// TL;DR Domain
type Domain struct {
	Name string `json:"name"`
	Slug string `json:"slug"`
}

// TL;DR language
type Language struct {
	Language   string  `json:"language"`
	Confidence float64 `json:"confidence"`
}

// TL;DR distribution channels
type DistributionChannels struct {
	LatestTldrs        bool `json:"latestTldrs"`
	LatestTldrsRSSFeed bool `json:"latestTldrsRSSFeed"`
}

// TL;DR creator information
type Creator struct {
	Id                 string `json:"id"`
	Username           string `json:"username"`
	UsernameForDisplay string `json:"usernameForDisplay"`
	IsAdmin            bool   `json:"isAdmin"`
	Deleted            bool   `json:"deleted"`
	TwitterHandle      string `json:"twitterHandle"`
}

// Struct reprsentation of a single TL;DR
type Tldr struct {
	Title                string               `json:"title"`
	Slug                 string               `json:"slug"`
	Permalink            string               `json:"permalink"`
	SummaryBullets       []string             `json:"summaryBullets"`
	ReadCount            int                  `json:"readCount"`
	OriginalUrl          string               `json:"originalUrl"`
	PossibleUrls         []string             `json:"possibleUrls"`
	Creator              Creator              `json:"creator"`
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
	Editors              []Creator            `json:"editors"`
	ThankedBy            []string             `json:"thankedBy"`
}

// Get the latest TL;DRs
func (t *TldrIo) Latest(number int, category string) (*[]Tldr, error) {
	var tldrs []Tldr
	v := url.Values{}

	if category != "" {
		v.Add("category", category)
	}

	err := callApi(t, "GET", fmt.Sprintf("tldrs/latest/%d", number), v.Encode(), &tldrs)
	if err != nil {
		return nil, err
	}
	return &tldrs, nil
}

// Search for a TL;DR by URL.
func (t *TldrIo) Search(u string) (*Tldr, error) {
	var tldr Tldr
	v := url.Values{}

	v.Set("url", u)

	err := callApi(t, "GET", "tldrs/search", v.Encode(), &tldr)
	if err != nil {
		return nil, err
	}
	return &tldr, nil

}



