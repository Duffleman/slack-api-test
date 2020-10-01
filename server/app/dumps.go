package app

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/kr/pretty"
)

func DumpRaw(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		HandleError(res, err)
		return
	}

	pretty.Println(string(body))

	HandleError(res, errors.New("debug_mode_on"))

	return
}

func DumpURLEncode(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		HandleError(res, err)
		return
	}

	pretty.Println(string(body))

	values, err := url.ParseQuery(string(body))
	if err != nil {
		HandleError(res, err)
		return
	}

	pretty.Println(values)

	HandleError(res, errors.New("debug_mode_on"))

	return
}

func DumpJSON(res http.ResponseWriter, req *http.Request) {
	var request map[string]interface{}

	err := json.NewDecoder(req.Body).Decode(&request)
	if err != nil {
		HandleError(res, err)
		return
	}

	pretty.Println(request)

	HandleError(res, errors.New("debug_mode_on"))

	return
}
