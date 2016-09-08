package model

type SystemJobs struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Created        string  `json:"created"`
		Description    string  `json:"description"`
		Elapsed        float64 `json:"elapsed"`
		ExtraVars      string  `json:"extra_vars"`
		Failed         bool    `json:"failed"`
		Finished       string  `json:"finished"`
		ID             int     `json:"id"`
		JobExplanation string  `json:"job_explanation"`
		JobType        string  `json:"job_type"`
		LaunchType     string  `json:"launch_type"`
		Modified       string  `json:"modified"`
		Name           string  `json:"name"`
		Related        struct {
			Cancel             string `json:"cancel"`
			Notifications      string `json:"notifications"`
			SystemJobTemplate  string `json:"system_job_template"`
			UnifiedJobTemplate string `json:"unified_job_template"`
		} `json:"related"`
		Started       string `json:"started"`
		Status        string `json:"status"`
		SummaryFields struct {
			UnifiedJobTemplate struct {
				Description    string `json:"description"`
				ID             int    `json:"id"`
				Name           string `json:"name"`
				UnifiedJobType string `json:"unified_job_type"`
			} `json:"unified_job_template"`
		} `json:"summary_fields"`
		SystemJobTemplate  int    `json:"system_job_template"`
		Type               string `json:"type"`
		UnifiedJobTemplate int    `json:"unified_job_template"`
		URL                string `json:"url"`
	} `json:"results"`
}
