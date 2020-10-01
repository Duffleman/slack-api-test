package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"serverdump/server/app"

	"github.com/kr/pretty"
)

var payload = map[string]interface{}{
	"blocks": []map[string]interface{}{
		{
			"type":     "section",
			"block_id": "select-my-shit",
			"text": map[string]interface{}{
				"type": "mrkdwn",
				"text": "Pick a shift you want to take",
			},
			"accessory": map[string]interface{}{
				"action_id": "text-my-shit",
				"type":      "static_select",
				"placeholder": map[string]interface{}{
					"type": "plain_text",
					"text": "Select an item",
				},
				"options": []map[string]interface{}{
					{
						"text": map[string]interface{}{
							"type": "plain_text",
							"text": "Option 1",
						},
						"value": "option-1",
					},
					{
						"text": map[string]interface{}{
							"type": "plain_text",
							"text": "Option 2",
						},
						"value": "option-2",
					},
				},
			},
		},
	},
}

func messageHandler(res http.ResponseWriter, req *http.Request) {
	err := safeHeaderCheck(req)
	if err != nil {
		app.HandleError(res, err)
		return
	}

	res.Header().Set("Content-Type", "application/json")

	json.NewEncoder(res).Encode(payload)
}

func interactionHandler(res http.ResponseWriter, req *http.Request) {
	err := safeHeaderCheck(req)
	if err != nil {
		app.HandleError(res, err)
		return
	}

	body, err := app.ParseInteractionRequest(req)
	if err != nil {
		app.HandleError(res, err)
		return
	}

	outgoing := fmt.Sprintf("You selected %s", body.Actions[0].SelectedOption.Value)

	httpClient := http.DefaultClient

	slackReturn, _ := json.Marshal(map[string]interface{}{
		"text": outgoing,
	})

	slackRes, err := httpClient.Post(body.ResponseURL, "application/json", bytes.NewBuffer(slackReturn))
	if err != nil {
		app.HandleError(res, err)
		return
	}

	slackResBody, err := ioutil.ReadAll(slackRes.Body)
	if err != nil {
		app.HandleError(res, err)
		return
	}

	pretty.Println(string(slackResBody))

	return
}

func safeHeaderCheck(req *http.Request) error {
	if v := req.Header.Get("content-type"); v != "" {
		if v != "application/x-www-form-urlencoded" {
			return errors.New("unhandled_content_type")
		}
	}

	return nil
}

func main() {
	http.HandleFunc("/message", messageHandler)
	http.HandleFunc("/interact", interactionHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func parseUserName(username string) string {
	names := strings.Split(username, ".")
	personalName := names[0]

	return strings.Title(personalName)
}
