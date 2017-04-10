package main

type Request struct {
	Words       []string    `json:"words"`
	NumHelpers  int         `json:"num_helpers"`
}

type JsonResponseBody struct {
	Pairs       []WordAndDefinition     `json:"data"`
	Time        string                  `json:"time"`
}

type WordAndDefinition struct {
	Word        string  `json:"word"`
	Definition  string  `json:"definition"`
}

type UrbanDictionaryInnerComponent struct {
	Definition  string  `json:"definition"`
}

type UrbanDictionaryResponse struct {
	List        []UrbanDictionaryInnerComponent     `json:"list"`
}