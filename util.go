package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func fetchJson(url string, data interface{}) error {
	if resp, err := http.Get(url); err != nil {
		return err
	} else {
		if bytes, err := ioutil.ReadAll(resp.Body); err != nil {
			return err
		} else {
			return json.Unmarshal(bytes, data)
		}
	}
}
