package towerapi

type JobEvents struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []JobEvent `json:"results"`
}

type JobEvent struct {
	Changed      bool     `json:"changed"`
	Counter      int      `json:"counter"`
	Created      string   `json:"created"`
	Event        string   `json:"event"`
	EventData    struct{} `json:"event_data"`
	EventDisplay string   `json:"event_display"`
	EventLevel   int      `json:"event_level"`
	Failed       bool     `json:"failed"`
	Host         string   `json:"host"`
	HostName     string   `json:"host_name"`
	ID           int      `json:"id"`
	Job          int      `json:"job"`
	Modified     string   `json:"modified"`
	Parent       string   `json:"parent"`
	Play         string   `json:"play"`
	Related      struct {
		Children string `json:"children"`
		Job      string `json:"job"`
	} `json:"related"`
	Role          string `json:"role"`
	SummaryFields struct {
		Job struct {
			Description     string `json:"description"`
			Failed          bool   `json:"failed"`
			ID              int    `json:"id"`
			JobTemplateID   int    `json:"job_template_id"`
			JobTemplateName string `json:"job_template_name"`
			Name            string `json:"name"`
			Status          string `json:"status"`
		} `json:"job"`
		Role struct{} `json:"role"`
	} `json:"summary_fields"`
	Task string `json:"task"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
