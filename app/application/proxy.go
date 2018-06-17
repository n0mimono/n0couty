package application

import (
	"io"
	"io/ioutil"
	"net/http"
)

func proxyApiGet(path string, method string) string {
	uri := path

	req, err := http.NewRequest(method, uri, nil)
	if err != nil {
		return `{ error: "Error" }`
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return `{ error: "Error" }`
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return `{ error: "Error" }`
	}

	return string(bytes)
}

func proxyApiPost(path string, method string, body io.Reader) string {
	uri := path

	req, err := http.NewRequest(method, uri, body)
	if err != nil {
		return `{ error: "Error" }`
	}
	req.Header.Set("Content-Type", "application/json")

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return `{ error: "Error" }`
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return `{ error: "Error" }`
	}

	return string(bytes)
}
