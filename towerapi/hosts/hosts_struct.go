package hosts

type Request struct {
	ID          string `json:"-"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Inventory   *int   `json:"inventory"`
	Enabled     bool   `json:"enabled"`
	InstanceID  string `json:"instance_id"`
	Variables   string `json:"variables" yaml:"variables"`
}

type Hosts struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Host `json:"results"`
}

type Host struct {
	Created             string `json:"created"`
	Description         string `json:"description"`
	Enabled             bool   `json:"enabled"`
	HasActiveFailures   bool   `json:"has_active_failures"`
	HasInventorySources bool   `json:"has_inventory_sources"`
	ID                  int    `json:"id"`
	InstanceID          string `json:"instance_id"`
	Inventory           int    `json:"inventory"`
	LastJob             int    `json:"last_job"`
	LastJobHostSummary  int    `json:"last_job_host_summary"`
	Modified            string `json:"modified"`
	Name                string `json:"name"`
	Related             struct {
		ActivityStream     string `json:"activity_stream"`
		AdHocCommandEvents string `json:"ad_hoc_command_events"`
		AdHocCommands      string `json:"ad_hoc_commands"`
		AllGroups          string `json:"all_groups"`
		CreatedBy          string `json:"created_by"`
		FactVersions       string `json:"fact_versions"`
		Groups             string `json:"groups"`
		Inventory          string `json:"inventory"`
		InventorySources   string `json:"inventory_sources"`
		JobEvents          string `json:"job_events"`
		JobHostSummaries   string `json:"job_host_summaries"`
		LastJob            string `json:"last_job"`
		LastJobHostSummary string `json:"last_job_host_summary"`
		ModifiedBy         string `json:"modified_by"`
		VariableData       string `json:"variable_data"`
	} `json:"related"`
	SummaryFields struct {
		CreatedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"created_by"`
		Inventory struct {
			Description                  string `json:"description"`
			GroupsWithActiveFailures     int    `json:"groups_with_active_failures"`
			HasActiveFailures            bool   `json:"has_active_failures"`
			HasInventorySources          bool   `json:"has_inventory_sources"`
			HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
			ID                           int    `json:"id"`
			InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
			Name                         string `json:"name"`
			TotalGroups                  int    `json:"total_groups"`
			TotalHosts                   int    `json:"total_hosts"`
			TotalInventorySources        int    `json:"total_inventory_sources"`
		} `json:"inventory"`
		LastJob struct {
			Description     string `json:"description"`
			Failed          bool   `json:"failed"`
			Finished        string `json:"finished"`
			ID              int    `json:"id"`
			JobTemplateID   int    `json:"job_template_id"`
			JobTemplateName string `json:"job_template_name"`
			Name            string `json:"name"`
			Status          string `json:"status"`
		} `json:"last_job"`
		LastJobHostSummary struct {
			Failed bool `json:"failed"`
			ID     int  `json:"id"`
		} `json:"last_job_host_summary"`
		ModifiedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"modified_by"`
		RecentJobs []struct {
			Finished string `json:"finished"`
			ID       int    `json:"id"`
			Name     string `json:"name"`
			Status   string `json:"status"`
		} `json:"recent_jobs"`
	} `json:"summary_fields"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Variables string `json:"variables"`
}
