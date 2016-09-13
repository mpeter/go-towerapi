package towerapi

type Roles struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Role `json:"results"`
}

type Role struct {
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
}
