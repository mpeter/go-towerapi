package credentials

type Request struct {
	ID                string `json:"-"`
	Name              string `json:"name"`
	Description       string `json:"description"`
	Kind              string `json:"kind"`
	Host              string `json:"host"`
	Username          string `json:"username"`
	Password          string `json:"password"`
	SecurityToken     string `json:"security_token"`
	Project           string `json:"project" mapstructure:"project_id"`
	Domain            string `json:"domain"`
	SSHKeyData        string `json:"ssh_key_data"`
	SSHKeyUnlock      string `json:"ssh_key_unlock"`
	Organization      *int   `json:"organization" mapstructure:"organization_id"`
	BecomeMethod      string `json:"become_method"`
	BecomeUsername    string `json:"become_username"`
	BecomePassword    string `json:"become_password"`
	VaultPassword     string `json:"vault_password"`
	Subscription      string `json:"subscription"`
	Tenant            string `json:"tenant"`
	Secret            string `json:"secret"`
	Client            string `json:"client"`
	Authorize         bool   `json:"authorize"`
	AuthorizePassword string `json:"authorize_password"`
}

type Credentials struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []Credential `json:"results"`
}

type Credential struct {
	Authorize         bool   `json:"authorize"`
	AuthorizePassword string `json:"authorize_password"`
	BecomeMethod      string `json:"become_method"`
	BecomePassword    string `json:"become_password"`
	BecomeUsername    string `json:"become_username"`
	Client            string `json:"client"`
	Cloud             bool   `json:"cloud"`
	Created           string `json:"created"`
	Description       string `json:"description"`
	Domain            string `json:"domain"`
	Host              string `json:"host"`
	ID                int    `json:"id"`
	Kind              string `json:"kind"`
	Modified          string `json:"modified"`
	Name              string `json:"name"`
	Organization      *int   `json:"organization"`
	Password          string `json:"password"`
	Project           string `json:"project"`
	Related           struct {
		AccessList     string `json:"access_list"`
		ActivityStream string `json:"activity_stream"`
		CreatedBy      string `json:"created_by"`
		ModifiedBy     string `json:"modified_by"`
		ObjectRoles    string `json:"object_roles"`
		Organization   string `json:"organization"`
		OwnerTeams     string `json:"owner_teams"`
		OwnerUsers     string `json:"owner_users"`
	} `json:"related"`
	Secret        string `json:"secret"`
	SecurityToken string `json:"security_token"`
	SSHKeyData    string `json:"ssh_key_data"`
	SSHKeyUnlock  string `json:"ssh_key_unlock"`
	Subscription  string `json:"subscription"`
	SummaryFields struct {
		CreatedBy struct {
			FirstName string `json:"first_name"`
			ID        *int   `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"created_by"`
		Host       struct{} `json:"host"`
		ModifiedBy struct {
			FirstName string `json:"first_name"`
			ID        *int   `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"modified_by"`
		ObjectRoles struct {
			AdminRole struct {
				Description string `json:"description"`
				ID          *int   `json:"id"`
				Name        string `json:"name"`
			} `json:"admin_role"`
			ReadRole struct {
				Description string `json:"description"`
				ID          *int   `json:"id"`
				Name        string `json:"name"`
			} `json:"read_role"`
			UseRole struct {
				Description string `json:"description"`
				ID          *int   `json:"id"`
				Name        string `json:"name"`
			} `json:"use_role"`
		} `json:"object_roles"`
		Organization struct {
			Description string `json:"description"`
			ID          *int   `json:"id"`
			Name        string `json:"name"`
		} `json:"organization"`
		Owners []struct {
			Description string `json:"description"`
			ID          *int   `json:"id"`
			Name        string `json:"name"`
			Type        string `json:"type"`
			URL         string `json:"url"`
		} `json:"owners"`
		Project struct{} `json:"project"`
	} `json:"summary_fields"`
	Tenant        string `json:"tenant"`
	Type          string `json:"type"`
	URL           string `json:"url"`
	Username      string `json:"username"`
	VaultPassword string `json:"vault_password"`
}
