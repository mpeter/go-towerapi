package model

type Roles struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Description string `json:"description"`
		ID          int    `json:"id"`
		Name        string `json:"name"`
		Related     struct {
			Teams string `json:"teams"`
			Users string `json:"users"`
		} `json:"related"`
		SummaryFields struct{} `json:"summary_fields"`
		Type          string   `json:"type"`
		URL           string   `json:"url"`
	} `json:"results"`
}
