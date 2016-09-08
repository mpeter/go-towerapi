package model

type Inventories struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Created                      string `json:"created"`
		Description                  string `json:"description"`
		GroupsWithActiveFailures     int    `json:"groups_with_active_failures"`
		HasActiveFailures            bool   `json:"has_active_failures"`
		HasInventorySources          bool   `json:"has_inventory_sources"`
		HostsWithActiveFailures      int    `json:"hosts_with_active_failures"`
		ID                           int    `json:"id"`
		InventorySourcesWithFailures int    `json:"inventory_sources_with_failures"`
		Modified                     string `json:"modified"`
		Name                         string `json:"name"`
		Organization                 int    `json:"organization"`
		Related                      struct {
			AccessList       string `json:"access_list"`
			ActivityStream   string `json:"activity_stream"`
			AdHocCommands    string `json:"ad_hoc_commands"`
			CreatedBy        string `json:"created_by"`
			Groups           string `json:"groups"`
			Hosts            string `json:"hosts"`
			InventorySources string `json:"inventory_sources"`
			JobTemplates     string `json:"job_templates"`
			ObjectRoles      string `json:"object_roles"`
			Organization     string `json:"organization"`
			RootGroups       string `json:"root_groups"`
			ScanJobTemplates string `json:"scan_job_templates"`
			Script           string `json:"script"`
			Tree             string `json:"tree"`
			VariableData     string `json:"variable_data"`
		} `json:"related"`
		SummaryFields struct {
			CreatedBy struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"created_by"`
			ObjectRoles struct {
				AdhocRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"adhoc_role"`
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
		TotalGroups           int    `json:"total_groups"`
		TotalHosts            int    `json:"total_hosts"`
		TotalInventorySources int    `json:"total_inventory_sources"`
		Type                  string `json:"type"`
		URL                   string `json:"url"`
		Variables             string `json:"variables"`
	} `json:"results"`
}
