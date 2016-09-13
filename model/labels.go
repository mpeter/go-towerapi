package towerapi

type Labels struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous string  `json:"previous"`
	Results  []Label `json:"results"`
}

type Label struct {
	Created      string `json:"created"`
	ID           int    `json:"id"`
	Modified     string `json:"modified"`
	Name         string `json:"name"`
	Organization int    `json:"organization"`
	Related      struct {
		CreatedBy    string `json:"created_by"`
		ModifiedBy   string `json:"modified_by"`
		Organization string `json:"organization"`
	} `json:"related"`
	SummaryFields struct {
		CreatedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"created_by"`
		ModifiedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"modified_by"`
		Organization struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"organization"`
	} `json:"summary_fields"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
