package app

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

func ParseInteractionRequest(req *http.Request) (*InteractionRequest, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	values, err := url.ParseQuery(string(body))
	if err != nil {
		return nil, err
	}

	payloadRaw := values.Get("payload")

	var payload *InteractionRequest

	err = json.Unmarshal([]byte(payloadRaw), &payload)
	if err != nil {
		return nil, err
	}

	return payload, nil
}
