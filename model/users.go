package towerapi

type Users struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []User `json:"results"`
}

type User struct {
	Auth            []interface{} `json:"auth"`
	Created         string        `json:"created"`
	Email           string        `json:"email"`
	ExternalAccount string        `json:"external_account"`
	FirstName       string        `json:"first_name"`
	ID              int           `json:"id"`
	IsSuperuser     bool          `json:"is_superuser"`
	IsSystemAuditor bool          `json:"is_system_auditor"`
	LastName        string        `json:"last_name"`
	LdapDn          string        `json:"ldap_dn"`
	Related         struct {
		AccessList           string `json:"access_list"`
		ActivityStream       string `json:"activity_stream"`
		AdminOfOrganizations string `json:"admin_of_organizations"`
		Credentials          string `json:"credentials"`
		Organizations        string `json:"organizations"`
		Projects             string `json:"projects"`
		Roles                string `json:"roles"`
		Teams                string `json:"teams"`
	} `json:"related"`
	Type     string `json:"type"`
	URL      string `json:"url"`
	Username string `json:"username"`
}
