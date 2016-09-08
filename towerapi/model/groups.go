package model

type Groups struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
		Created                  string `json:"created"`
		Description              string `json:"description"`
		GroupsWithActiveFailures int    `json:"groups_with_active_failures"`
		HasActiveFailures        bool   `json:"has_active_failures"`
		HasInventorySources      bool   `json:"has_inventory_sources"`
		HostsWithActiveFailures  int    `json:"hosts_with_active_failures"`
		ID                       int    `json:"id"`
		Inventory                int    `json:"inventory"`
		Modified                 string `json:"modified"`
		Name                     string `json:"name"`
		Related                  struct {
			ActivityStream    string `json:"activity_stream"`
			AdHocCommands     string `json:"ad_hoc_commands"`
			AllHosts          string `json:"all_hosts"`
			Children          string `json:"children"`
			CreatedBy         string `json:"created_by"`
			Hosts             string `json:"hosts"`
			Inventory         string `json:"inventory"`
			InventorySource   string `json:"inventory_source"`
			InventorySources  string `json:"inventory_sources"`
			JobEvents         string `json:"job_events"`
			JobHostSummaries  string `json:"job_host_summaries"`
			ModifiedBy        string `json:"modified_by"`
			PotentialChildren string `json:"potential_children"`
			VariableData      string `json:"variable_data"`
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
			InventorySource struct {
				Source string `json:"source"`
				Status string `json:"status"`
			} `json:"inventory_source"`
			ModifiedBy struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"modified_by"`
		} `json:"summary_fields"`
		TotalGroups int    `json:"total_groups"`
		TotalHosts  int    `json:"total_hosts"`
		Type        string `json:"type"`
		URL         string `json:"url"`
		Variables   string `json:"variables"`
	} `json:"results"`
}
