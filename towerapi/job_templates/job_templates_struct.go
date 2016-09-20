package job_templates

type Request struct {
	ID                    string    `json:"-"`
	Name                  string `json:"name"`
	Description           string `json:"description"`
	JobType               string `json:"job_type"`
	Inventory             *int   `json:"inventory"`
	Project               *int   `json:"project"`
	Playbook              string `json:"playbook"`
	Credential            *int   `json:"credential"`
	CloudCredential       *int   `json:"cloud_credential"`
	NetworkCredential     *int   `json:"network_credential"`
	Forks                 int    `json:"forks"`
	Limit                 string `json:"limit"`
	Verbosity             int    `json:"verbosity"`
	ExtraVars             string `json:"extra_vars"`
	JobTags               string `json:"job_tags"`
	ForceHandlers         bool   `json:"force_handlers"`
	SkipTags              string `json:"skip_tags"`
	StartAtTask           string `json:"start_at_task"`
	HostConfigKey         string `json:"host_config_key"`
	AskVariablesOnLaunch  bool   `json:"ask_variables_on_launch"`
	AskLimitOnLaunch      bool   `json:"ask_limit_on_launch"`
	AskTagsOnLaunch       bool   `json:"ask_tags_on_launch"`
	AskSkipTagsOnLaunch   bool   `json:"ask_skip_tags_on_launch"`
	AskJobTypeOnLaunch    bool   `json:"ask_job_type_on_launch"`
	AskCredentialOnLaunch bool   `json:"ask_credential_on_launch"`
	SurveyEnabled         bool   `json:"survey_enabled"`
	AllowSimultaneous     bool   `json:"allow_simultaneous"`
}

type JobTemplates struct {
	Count    int           `json:"count"`
	Next     string        `json:"next"`
	Previous string        `json:"previous"`
	Results  []JobTemplate `json:"results"`
}

type JobTemplate struct {
	AllowSimultaneous     bool   `json:"allow_simultaneous"`
	AskCredentialOnLaunch bool   `json:"ask_credential_on_launch"`
	AskInventoryOnLaunch  bool   `json:"ask_inventory_on_launch"`
	AskJobTypeOnLaunch    bool   `json:"ask_job_type_on_launch"`
	AskLimitOnLaunch      bool   `json:"ask_limit_on_launch"`
	AskSkipTagsOnLaunch   bool   `json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch       bool   `json:"ask_tags_on_launch"`
	AskVariablesOnLaunch  bool   `json:"ask_variables_on_launch"`
	BecomeEnabled         bool   `json:"become_enabled"`
	CloudCredential       *int   `json:"cloud_credential"`
	Created               string `json:"created"`
	Credential            *int   `json:"credential"`
	Description           string `json:"description"`
	ExtraVars             string `json:"extra_vars"`
	ForceHandlers         bool   `json:"force_handlers"`
	Forks                 int    `json:"forks"`
	HasSchedules          bool   `json:"has_schedules"`
	HostConfigKey         string `json:"host_config_key"`
	ID                    int    `json:"id"`
	Inventory             *int   `json:"inventory"`
	JobTags               string `json:"job_tags"`
	JobType               string `json:"job_type"`
	LastJobFailed         bool   `json:"last_job_failed"`
	LastJobRun            string `json:"last_job_run"`
	Limit                 string `json:"limit"`
	Modified              string `json:"modified"`
	Name                  string `json:"name"`
	NetworkCredential     *int   `json:"network_credential"`
	NextJobRun            string `json:"next_job_run"`
	Playbook              string `json:"playbook"`
	Project               *int   `json:"project"`
	Related               struct {
		AccessList                   string `json:"access_list"`
		ActivityStream               string `json:"activity_stream"`
		Callback                     string `json:"callback"`
		CloudCredential              string `json:"cloud_credential"`
		CreatedBy                    string `json:"created_by"`
		Credential                   string `json:"credential"`
		Inventory                    string `json:"inventory"`
		Jobs                         string `json:"jobs"`
		Labels                       string `json:"labels"`
		LastJob                      string `json:"last_job"`
		Launch                       string `json:"launch"`
		ModifiedBy                   string `json:"modified_by"`
		NetworkCredential            string `json:"network_credential"`
		NotificationTemplatesAny     string `json:"notification_templates_any"`
		NotificationTemplatesError   string `json:"notification_templates_error"`
		NotificationTemplatesSuccess string `json:"notification_templates_success"`
		ObjectRoles                  string `json:"object_roles"`
		Project                      string `json:"project"`
		Schedules                    string `json:"schedules"`
		SurveySpec                   string `json:"survey_spec"`
	} `json:"related"`
	SkipTags      string `json:"skip_tags"`
	StartAtTask   string `json:"start_at_task"`
	Status        string `json:"status"`
	SummaryFields struct {
		CanCopy         bool `json:"can_copy"`
		CanEdit         bool `json:"can_edit"`
		CloudCredential struct {
			Cloud       bool   `json:"cloud"`
			Description string `json:"description"`
			ID          int    `json:"id"`
			Kind        string `json:"kind"`
			Name        string `json:"name"`
		} `json:"cloud_credential"`
		CreatedBy struct {
			FirstName string `json:"first_name"`
			ID        int   `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"created_by"`
		Credential struct {
			Cloud       bool   `json:"cloud"`
			Description string `json:"description"`
			ID          int   `json:"id"`
			Kind        string `json:"kind"`
			Name        string `json:"name"`
		} `json:"credential"`
		Inventory struct {
			Description                  string `json:"description"`
			GroupsWithActiveFailures     int    `json:"groups_with_active_failures"`
			HasActiveFailures            bool   `json:"has_active_failures"`
			HasInventorySources          bool   `json:"has_inventory_sources"`
			HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
			ID                           int   `json:"id"`
			InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
			Name                         string `json:"name"`
			TotalGroups                  int    `json:"total_groups"`
			TotalHosts                   int    `json:"total_hosts"`
			TotalInventorySources        int    `json:"total_inventory_sources"`
		} `json:"inventory"`
		Labels struct {
			Count   int `json:"count"`
			Results []struct {
				ID   int   `json:"id"`
				Name string `json:"name"`
			} `json:"results"`
		} `json:"labels"`
		LastJob struct {
			Description string `json:"description"`
			Failed      bool   `json:"failed"`
			Finished    string `json:"finished"`
			ID          int   `json:"id"`
			Name        string `json:"name"`
			Status      string `json:"status"`
		} `json:"last_job"`
		LastUpdate struct {
			Description string `json:"description"`
			Failed      bool   `json:"failed"`
			ID          int   `json:"id"`
			Name        string `json:"name"`
			Status      string `json:"status"`
		} `json:"last_update"`
		ModifiedBy struct {
			FirstName string `json:"first_name"`
			ID        int   `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"modified_by"`
		NetworkCredential struct {
			Description string `json:"description"`
			ID          int   `json:"id"`
			Kind        string `json:"kind"`
			Name        string `json:"name"`
		} `json:"network_credential"`
		ObjectRoles struct {
			AdminRole struct {
				Description string `json:"description"`
				ID          int   `json:"id"`
				Name        string `json:"name"`
			} `json:"admin_role"`
			ExecuteRole struct {
				Description string `json:"description"`
				ID          int   `json:"id"`
				Name        string `json:"name"`
			} `json:"execute_role"`
			ReadRole struct {
				Description string `json:"description"`
				ID          int   `json:"id"`
				Name        string `json:"name"`
			} `json:"read_role"`
		} `json:"object_roles"`
		Project struct {
			Description string `json:"description"`
			ID          int   `json:"id"`
			Name        string `json:"name"`
			Status      string `json:"status"`
		} `json:"project"`
		RecentJobs []struct {
			Finished string `json:"finished"`
			ID       int   `json:"id"`
			Status   string `json:"status"`
		} `json:"recent_jobs"`
	} `json:"summary_fields"`
	SurveyEnabled bool   `json:"survey_enabled"`
	Type          string `json:"type"`
	URL           string `json:"url"`
	Verbosity     int    `json:"verbosity"`
}
