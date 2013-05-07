# tldriogo

A Go library for the tldr.io API.

[![Build Status](https://travis-ci.org/joshrotenberg/tldriogo.png?branch=master)](https://travis-ci.org/joshrotenberg/tldriogo)

## Overview

tldriogo is a client library for the [tl;dr](http://tldr.io/) API.

## Summary

```go

var tldrio *TldrIo
tldrio = NewTldrIo()

categories, err := tldrio.Categories()
// &[{Name:Tech News Slug:tech-news} {Name:World News Slug:world-news} ...]

// Get the three latest World News TL;DRs 
latest, err := tldrio.Latest(3, "world-news")

// Search for a TL;DR for a particular URL
result, err := tldrio.Search("http://i.want.a.summary")

// Search for a bunch of TL;DRs
results, err := tldrio.SearchBatch("http://i.want.a.summary", "http://and.another")

// Get the scoop on a TL;DR user
user, err := tldrio.User("fredjonez")

// See all of frejonez' TL;DRs
results, err := tldrio.

// SGet te scoop on a TL;DR user
ruer, err := tldrio.User("fredjonez")

// See all of frejonez' TL;DRs
results, err := tldrio.UserTldrs("fredjonez")
```

## Docs
See http://godoc.org/github.com/joshrotenberg/tldriogo and http://tldr.io/api-documentation for more information.


