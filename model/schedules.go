package towerapi

type Schedules struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous string     `json:"previous"`
	Results  []Schedule `json:"results"`
}

type Schedule struct {
	Created     string `json:"created"`
	Description string `json:"description"`
	Dtend       string `json:"dtend"`
	Dtstart     string `json:"dtstart"`
	Enabled     bool   `json:"enabled"`
	ExtraData   struct {
		Days string `json:"days"`
	} `json:"extra_data"`
	ID       int    `json:"id"`
	Modified string `json:"modified"`
	Name     string `json:"name"`
	NextRun  string `json:"next_run"`
	Related  struct {
		UnifiedJobTemplate string `json:"unified_job_template"`
		UnifiedJobs        string `json:"unified_jobs"`
	} `json:"related"`
	Rrule         string `json:"rrule"`
	SummaryFields struct {
		UnifiedJobTemplate struct {
			Description    string `json:"description"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			UnifiedJobType string `json:"unified_job_type"`
		} `json:"unified_job_template"`
	} `json:"summary_fields"`
	Type               string `json:"type"`
	UnifiedJobTemplate int    `json:"unified_job_template"`
	URL                string `json:"url"`
}
