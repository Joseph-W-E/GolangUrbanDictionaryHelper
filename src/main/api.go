package main

import (
	"net/http"
	"time"
	"io/ioutil"
	"encoding/json"
	"log"
)

func SendRequest(word string) string {
	log.Println("Attempting request for word:", word)

	// setup the request
	request, error := http.NewRequest("GET", "https://mashape-community-urban-dictionary.p.mashape.com/define?term=" + word, nil)
	if error != nil {
		return ""
	}
	request.Header.Set("X-Mashape-Key", getKey())
	request.Header.Set("Accept", "text/plain")

	// setup the client (with a timeout)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(request)
	if err != nil {
		return ""
	}

	// read the response into an UrbanDictionaryResponse struct
	var urbanDictionaryResponse UrbanDictionaryResponse
	content, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(content, &urbanDictionaryResponse)
	if err != nil {
		return ""
	}

	// return the first element in the list of definitions, if any
	if len(urbanDictionaryResponse.List) > 0 {
		return urbanDictionaryResponse.List[0].Definition
	} else {
		return ""
	}
}

/* Totally not an API key hard-coded (it's public, chill) */

func getKey() string {
	return ""
}
