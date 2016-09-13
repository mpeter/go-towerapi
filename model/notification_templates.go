package towerapi

type NotificationTemplates struct {
	Count    int                    `json:"count"`
	Next     string                 `json:"next"`
	Previous string                 `json:"previous"`
	Results  []NotificationTemplate `json:"results"`
}

type NotificationTemplate struct {
	Created                   string `json:"created"`
	Description               string `json:"description"`
	ID                        int    `json:"id"`
	Modified                  string `json:"modified"`
	Name                      string `json:"name"`
	NotificationConfiguration struct {
		Headers struct {
			X_Ansible    string `json:"X-Ansible"`
			X_Auth_Token string `json:"X-Auth-Token"`
		} `json:"headers"`
		URL string `json:"url"`
	} `json:"notification_configuration"`
	NotificationType string `json:"notification_type"`
	Organization     int    `json:"organization"`
	Related          struct {
		CreatedBy     string `json:"created_by"`
		ModifiedBy    string `json:"modified_by"`
		Notifications string `json:"notifications"`
		Organization  string `json:"organization"`
		Test          string `json:"test"`
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
		Organization struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"organization"`
		RecentNotifications []struct {
			Created string `json:"created"`
			ID      int    `json:"id"`
			Status  string `json:"status"`
		} `json:"recent_notifications"`
	} `json:"summary_fields"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
