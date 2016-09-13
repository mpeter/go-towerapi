package organizations

type Organizations struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []Organization `json:"results"`
}

type Organization struct {
	Created     string `json:"created"`
	Description string `json:"description"`
	ID          int    `json:"id"`
	Modified    string `json:"modified"`
	Name        string `json:"name"`
	Related     struct {
		AccessList                   string `json:"access_list"`
		ActivityStream               string `json:"activity_stream"`
		Admins                       string `json:"admins"`
		CreatedBy                    string `json:"created_by"`
		Credentials                  string `json:"credentials"`
		Inventories                  string `json:"inventories"`
		ModifiedBy                   string `json:"modified_by"`
		NotificationTemplates        string `json:"notification_templates"`
		NotificationTemplatesAny     string `json:"notification_templates_any"`
		NotificationTemplatesError   string `json:"notification_templates_error"`
		NotificationTemplatesSuccess string `json:"notification_templates_success"`
		ObjectRoles                  string `json:"object_roles"`
		Projects                     string `json:"projects"`
		Teams                        string `json:"teams"`
		Users                        string `json:"users"`
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
			AuditorRole struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
				Name        string `json:"name"`
			} `json:"auditor_role"`
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
		RelatedFieldCounts struct {
			Admins       int `json:"admins"`
			Inventories  int `json:"inventories"`
			JobTemplates int `json:"job_templates"`
			Projects     int `json:"projects"`
			Teams        int `json:"teams"`
			Users        int `json:"users"`
		} `json:"related_field_counts"`
	} `json:"summary_fields"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
