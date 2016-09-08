package model

type Config struct {
	AnalyticsStatus string `json:"analytics_status"`
	AnsibleVersion  string `json:"ansible_version"`
	Eula            string `json:"eula"`
	LicenseInfo     struct {
		AvailableInstances int    `json:"available_instances"`
		CompanyName        string `json:"company_name"`
		Compliant          bool   `json:"compliant"`
		ContactEmail       string `json:"contact_email"`
		ContactName        string `json:"contact_name"`
		CurrentInstances   int    `json:"current_instances"`
		DateExpired        bool   `json:"date_expired"`
		DateWarning        bool   `json:"date_warning"`
		DeploymentID       string `json:"deployment_id"`
		Features           struct {
			ActivityStreams       bool `json:"activity_streams"`
			EnterpriseAuth        bool `json:"enterprise_auth"`
			Ha                    bool `json:"ha"`
			Ldap                  bool `json:"ldap"`
			MultipleOrganizations bool `json:"multiple_organizations"`
			Rebranding            bool `json:"rebranding"`
			Surveys               bool `json:"surveys"`
			SystemTracking        bool `json:"system_tracking"`
		} `json:"features"`
		FreeInstances        int    `json:"free_instances"`
		GracePeriodRemaining int    `json:"grace_period_remaining"`
		Hostname             string `json:"hostname"`
		InstanceCount        int    `json:"instance_count"`
		LicenseDate          int    `json:"license_date"`
		LicenseKey           string `json:"license_key"`
		LicenseType          string `json:"license_type"`
		SubscriptionName     string `json:"subscription_name"`
		TimeRemaining        int    `json:"time_remaining"`
		ValidKey             bool   `json:"valid_key"`
	} `json:"license_info"`
	ProjectBaseDir    string        `json:"project_base_dir"`
	ProjectLocalPaths []interface{} `json:"project_local_paths"`
	TimeZone          string        `json:"time_zone"`
	Version           string        `json:"version"`
}
