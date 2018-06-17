package infrastructure

import (
	"io/ioutil"
	"net/http"
	"net/url"
)

func get(path string, values url.Values) (string, error) {
	uri := path
	if len(values) > 0 {
		uri = path + "?" + values.Encode()
	}

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return "", err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
