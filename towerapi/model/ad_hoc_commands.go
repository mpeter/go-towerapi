package model

type AdHocCommands struct {
	Count    int           `json:"count"`
	Next     interface{}   `json:"next"`
	Previous interface{}   `json:"previous"`
	Results  []interface{} `json:"results"`
}
