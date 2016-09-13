package towerapi

type InventorySources struct {
	Count    int               `json:"count"`
	Next     string            `json:"next"`
	Previous string            `json:"previous"`
	Results  []InventorySource `json:"results"`
}

type InventorySource struct {
	Created          string `json:"created"`
	Credential       string `json:"credential"`
	Description      string `json:"description"`
	Group            int    `json:"group"`
	GroupBy          string `json:"group_by"`
	HasSchedules     bool   `json:"has_schedules"`
	ID               int    `json:"id"`
	InstanceFilters  string `json:"instance_filters"`
	Inventory        int    `json:"inventory"`
	LastJobFailed    bool   `json:"last_job_failed"`
	LastJobRun       string `json:"last_job_run"`
	LastUpdateFailed bool   `json:"last_update_failed"`
	LastUpdated      string `json:"last_updated"`
	Modified         string `json:"modified"`
	Name             string `json:"name"`
	NextJobRun       string `json:"next_job_run"`
	Overwrite        bool   `json:"overwrite"`
	OverwriteVars    bool   `json:"overwrite_vars"`
	Related          struct {
		ActivityStream               string `json:"activity_stream"`
		CreatedBy                    string `json:"created_by"`
		Group                        string `json:"group"`
		Groups                       string `json:"groups"`
		Hosts                        string `json:"hosts"`
		Inventory                    string `json:"inventory"`
		InventoryUpdates             string `json:"inventory_updates"`
		ModifiedBy                   string `json:"modified_by"`
		NotificationTemplatesAny     string `json:"notification_templates_any"`
		NotificationTemplatesError   string `json:"notification_templates_error"`
		NotificationTemplatesSuccess string `json:"notification_templates_success"`
		Schedules                    string `json:"schedules"`
		Update                       string `json:"update"`
	} `json:"related"`
	Source        string `json:"source"`
	SourcePath    string `json:"source_path"`
	SourceRegions string `json:"source_regions"`
	SourceScript  string `json:"source_script"`
	SourceVars    string `json:"source_vars"`
	Status        string `json:"status"`
	SummaryFields struct {
		CreatedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"created_by"`
		Group struct {
			Description              string `json:"description"`
			GroupsWithActiveFailures int    `json:"groups_with_active_failures"`
			HasActiveFailures        bool   `json:"has_active_failures"`
			HasInventorySources      bool   `json:"has_inventory_sources"`
			HostsWithActiveFailures  int    `json:"hosts_with_active_failures"`
			ID                       int    `json:"id"`
			Name                     string `json:"name"`
			TotalGroups              int    `json:"total_groups"`
			TotalHosts               int    `json:"total_hosts"`
		} `json:"group"`
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
		ModifiedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"modified_by"`
	} `json:"summary_fields"`
	Type               string `json:"type"`
	UpdateCacheTimeout int    `json:"update_cache_timeout"`
	UpdateOnLaunch     bool   `json:"update_on_launch"`
	URL                string `json:"url"`
}
