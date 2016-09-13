package towerapi

type Teams struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Team `json:"results"`
}

type Team struct {
	Created      string `json:"created"`
	Description  string `json:"description"`
	ID           int    `json:"id"`
	Modified     string `json:"modified"`
	Name         string `json:"name"`
	Organization int    `json:"organization"`
	Related      struct {
		AccessList     string `json:"access_list"`
		ActivityStream string `json:"activity_stream"`
		CreatedBy      string `json:"created_by"`
		Credentials    string `json:"credentials"`
		ModifiedBy     string `json:"modified_by"`
		ObjectRoles    string `json:"object_roles"`
		Organization   string `json:"organization"`
		Projects       string `json:"projects"`
		Roles          string `json:"roles"`
		Users          string `json:"users"`
	} `json:"related"`
	SummaryFields struct {
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
			MemberRole struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
			} `json:"member_role"`
			ReadRole struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
			} `json:"read_role"`
		} `json:"object_roles"`
		Organization struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"organization"`
	} `json:"summary_fields"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
