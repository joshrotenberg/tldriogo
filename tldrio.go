package tldrio

import (
	"encoding/json"
	"errors"
	"fmt"
	//	"io"
	"io/ioutil"
	"net/http"
)

const (
	TldrApiUrl = "https://api.tldr.io"
)

type Tldr struct {
	httpClient *http.Client
}

func NewTldr() *Tldr {
	client := &http.Client{}
	return &Tldr{httpClient: client}

}

func callApi(t *Tldr, method, uri, params string, result interface{}) error {
	url := fmt.Sprintf("%s/%s?%s", TldrApiUrl, uri, params)
	request, err := http.NewRequest(method, url, nil)
	if err != nil {
		return err
	}

	response, err := t.httpClient.Do(request)
	defer response.Body.Close()
	if err != nil {
		return err
	}

	if response.StatusCode != http.StatusOK {
		err = errors.New(fmt.Sprintf("API responded with %d", response.StatusCode))
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

