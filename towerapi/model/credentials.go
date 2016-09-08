package model

type Credentials struct {
	Count    int         `json:"count"`
	Next     interface{} `json:"next"`
	Previous interface{} `json:"previous"`
	Results  []struct {
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
		Password          string `json:"password"`
		Project           string `json:"project"`
		Related           struct {
			AccessList     string `json:"access_list"`
			ActivityStream string `json:"activity_stream"`
			CreatedBy      string `json:"created_by"`
			ModifiedBy     string `json:"modified_by"`
			ObjectRoles    string `json:"object_roles"`
			OwnerTeams     string `json:"owner_teams"`
			OwnerUsers     string `json:"owner_users"`
			User           string `json:"user"`
		} `json:"related"`
		Secret        string `json:"secret"`
		SecurityToken string `json:"security_token"`
		SSHKeyData    string `json:"ssh_key_data"`
		SSHKeyUnlock  string `json:"ssh_key_unlock"`
		Subscription  string `json:"subscription"`
		SummaryFields struct {
			CreatedBy struct {
				FirstName string `json:"first_name"`
				ID        int    `json:"id"`
				LastName  string `json:"last_name"`
				Username  string `json:"username"`
			} `json:"created_by"`
			Host       struct{} `json:"host"`
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
				ReadRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"read_role"`
				UseRole struct {
					Description string `json:"description"`
					ID          int    `json:"id"`
					Name        string `json:"name"`
				} `json:"use_role"`
			} `json:"object_roles"`
			Owners []struct {
				Description string `json:"description"`
				ID          int    `json:"id"`
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
	} `json:"results"`
}
