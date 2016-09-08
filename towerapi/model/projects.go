package model

type Projects struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Created          string      `json:"created"`
		Credential       interface{} `json:"credential"`
		Description      string      `json:"description"`
		HasSchedules     bool        `json:"has_schedules"`
		ID               int         `json:"id"`
		LastJobFailed    bool        `json:"last_job_failed"`
		LastJobRun       interface{} `json:"last_job_run"`
		LastUpdateFailed bool        `json:"last_update_failed"`
		LastUpdated      interface{} `json:"last_updated"`
		LocalPath        string      `json:"local_path"`
		Modified         string      `json:"modified"`
		Name             string      `json:"name"`
		NextJobRun       interface{} `json:"next_job_run"`
		Organization     int         `json:"organization"`
		Related          struct {
			AccessList                   string `json:"access_list"`
			ActivityStream               string `json:"activity_stream"`
			CreatedBy                    string `json:"created_by"`
			ModifiedBy                   string `json:"modified_by"`
			NotificationTemplatesAny     string `json:"notification_templates_any"`
			NotificationTemplatesError   string `json:"notification_templates_error"`
			NotificationTemplatesSuccess string `json:"notification_templates_success"`
			ObjectRoles                  string `json:"object_roles"`
			Organization                 string `json:"organization"`
			Playbooks                    string `json:"playbooks"`
			ProjectUpdates               string `json:"project_updates"`
			Schedules                    string `json:"schedules"`
			Teams                        string `json:"teams"`
			Update                       string `json:"update"`
		} `json:"related"`
		ScmBranch             string `json:"scm_branch"`
		ScmClean              bool   `json:"scm_clean"`
		ScmDeleteOnNextUpdate bool   `json:"scm_delete_on_next_update"`
		ScmDeleteOnUpdate     bool   `json:"scm_delete_on_update"`
		ScmType               string `json:"scm_type"`
		ScmUpdateCacheTimeout int    `json:"scm_update_cache_timeout"`
		ScmUpdateOnLaunch     bool   `json:"scm_update_on_launch"`
		ScmURL                string `json:"scm_url"`
		Status                string `json:"status"`
		SummaryFields         struct {
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
			ObjectRoles struct {
				AdminRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"admin_role"`
				ReadRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"read_role"`
				UpdateRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"update_role"`
				UseRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"use_role"`
			} `json:"object_roles"`
			Organization struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
			} `json:"organization"`
		} `json:"summary_fields"`
		Type string `json:"type"`
		URL  string `json:"url"`
	} `json:"results"`
}
