// Package tldrio provides a simple wrapper around the tldr.io API.
package tldrio

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	TldrApiUrl = "https://api.tldr.io"
)

type TldrIo struct {
	httpClient *http.Client
}

// NewTldrIo creates a new handle for making API calls.
func NewTldrIo() *TldrIo {
	client := &http.Client{}
	return &TldrIo{httpClient: client}

}

func callApi(t *TldrIo, method, uri, params string, postData []byte, result interface{}) error {
	url := fmt.Sprintf("%s/%s?%s", TldrApiUrl, uri, params)
	var (
		err     error
		request *http.Request
	)
	switch method {
	case "GET":
		request, err = http.NewRequest(method, url, nil)
		if err != nil {
			return err
		}
	case "POST":
		request, err = http.NewRequest(method, url, bytes.NewReader(postData))
		if err != nil {
			return err
		}
		request.ContentLength = int64(len(postData))
		request.Header.Set("Content-Type", "application/json")
	default:
		err = errors.New(fmt.Sprintf("Method %s not supported", method))
		return err
	}

	request.Header.Add("User-Agent", "tldriogo/0.1")

	response, err := t.httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("API responded with %d", response.StatusCode))
		return err
	}

	if response.Header.Get("Content-Type") == "application/json; charset=utf-8" {
		var js []byte
		js, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		} else {
			err = json.Unmarshal(js, &result)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
