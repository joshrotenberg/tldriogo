# tldriogo

A Go library for the tldr.io API.

## Overview

tldriogo is a client library for the [tl;dr](http://tldr.io/) API.

## Docs

``go

var tldrio *TldrIo
tldrio = NewTldrIo()

categories, err := tldrio.Categories()
// &[{Name:Tech News Slug:tech-news} {Name:World News Slug:world-news} ...]

// Get the three latest World News TL;DRs 
latest, err := tldrio.Latest(3, "world-news")

// Search for a TL;DR for a particular URL
result, err := tldrio.Search("http://i.want.a.summary")
```

See http://godoc.org/github.com/joshrotenberg/tldriogo and http://tldr.io/api-documentation for more information.
