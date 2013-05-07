package tldrio

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

// TL;DR Domain
type Domain struct {
	Name string
	Slug string
}

// TL;DR language
type Language struct {
	Language   string
	Confidence float64
}

// TL;DR distribution channels
type DistributionChannels struct {
	LatestTldrs        bool
	LatestTldrsRSSFeed bool
}

// TL;DR creator information
type Creator struct {
	Id                 string
	Username           string
	UsernameForDisplay string
	IsAdmin            bool
	Deleted            bool
	TwitterHandle      string
}

// Struct reprsentation of a single TL;DR
type Tldr struct {
	Title                string
	Slug                 string
	Permalink            string
	SummaryBullets       []string
	ReadCount            int
	OriginalUrl          string
	PossibleUrls         []string
	Creator              Creator
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
	Editors              []Creator
	ThankedBy            []string
}

// Latest gets the latest TL;DRs.
func (t *TldrIo) Latest(number int, category string) (*[]Tldr, error) {
	var tldrs []Tldr
	v := url.Values{}

	if category != "" {
		v.Add("category", category)
	}

	err := callApi(t, "GET", fmt.Sprintf("tldrs/latest/%d", number), v.Encode(), nil, &tldrs)
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

	err := callApi(t, "GET", "tldrs/search", v.Encode(), nil, &tldr)
	if err != nil {
		return nil, err
	}
	return &tldr, nil

}

// SearchBatch searches for multiple TL;DRs.
func (t *TldrIo) SearchBatch(urls ...string) (*[]Tldr, error) {
	b, err := json.Marshal(struct {
		Batch []string `json:"batch"`
	}{urls})
	if err != nil {
		return nil, err
	}

	data := struct {
		Tldrs []Tldr
	}{}

	err = callApi(t, "POST", "tldrs/searchBatch", "", b, &data)
	if err != nil {
		return nil, err
	}
	return &data.Tldrs, nil
}
