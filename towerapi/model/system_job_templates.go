package model

type SystemJobTemplates struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Created       string `json:"created"`
		Description   string `json:"description"`
		HasSchedules  bool   `json:"has_schedules"`
		ID            int    `json:"id"`
		JobType       string `json:"job_type"`
		LastJobFailed bool   `json:"last_job_failed"`
		LastJobRun    string `json:"last_job_run"`
		Modified      string `json:"modified"`
		Name          string `json:"name"`
		NextJobRun    string `json:"next_job_run"`
		Related       struct {
			Jobs                         string `json:"jobs"`
			LastJob                      string `json:"last_job"`
			Launch                       string `json:"launch"`
			NextSchedule                 string `json:"next_schedule"`
			NotificationTemplatesAny     string `json:"notification_templates_any"`
			NotificationTemplatesError   string `json:"notification_templates_error"`
			NotificationTemplatesSuccess string `json:"notification_templates_success"`
			Schedules                    string `json:"schedules"`
		} `json:"related"`
		Status        string `json:"status"`
		SummaryFields struct {
			LastJob struct {
				Description string `json:"description"`
				Failed      bool   `json:"failed"`
				Finished    string `json:"finished"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Status      string `json:"status"`
			} `json:"last_job"`
			LastUpdate struct {
				Description string `json:"description"`
				Failed      bool   `json:"failed"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
				Status      string `json:"status"`
			} `json:"last_update"`
		} `json:"summary_fields"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"results"`
}
