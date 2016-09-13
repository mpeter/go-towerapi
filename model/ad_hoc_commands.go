package towerapi

type AdHocCommands struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []AdHocCommand `json:"results"`
}

type AdHocCommand struct {
	BecomeEnabled bool    `json:"become_enabled"`
	Created       string  `json:"created"`
	Credential    int     `json:"credential"`
	Elapsed       float64 `json:"elapsed"`
	ExtraVars     string  `json:"extra_vars"`
	Failed        bool    `json:"failed"`
	Finished      string  `json:"finished"`
	Forks         int     `json:"forks"`
	ID            int     `json:"id"`
	Inventory     int     `json:"inventory"`
	JobArgs       string  `json:"job_args"`
	JobCwd        string  `json:"job_cwd"`
	JobEnv        struct {
		ADHOCCOMMANDID                string `json:"AD_HOC_COMMAND_ID"`
		ANSIBLECALLBACKPLUGINS        string `json:"ANSIBLE_CALLBACK_PLUGINS"`
		ANSIBLEFORCECOLOR             string `json:"ANSIBLE_FORCE_COLOR"`
		ANSIBLEHOSTKEYCHECKING        string `json:"ANSIBLE_HOST_KEY_CHECKING"`
		ANSIBLELOADCALLBACKPLUGINS    string `json:"ANSIBLE_LOAD_CALLBACK_PLUGINS"`
		ANSIBLEPARAMIKORECORDHOSTKEYS string `json:"ANSIBLE_PARAMIKO_RECORD_HOST_KEYS"`
		ANSIBLESFTPBATCHMODE          string `json:"ANSIBLE_SFTP_BATCH_MODE"`
		ANSIBLESSHARGS                string `json:"ANSIBLE_SSH_ARGS"`
		ANSIBLEUSEVENV                string `json:"ANSIBLE_USE_VENV"`
		ANSIBLEVENVPATH               string `json:"ANSIBLE_VENV_PATH"`
		CALLBACKCONSUMERPORT          string `json:"CALLBACK_CONSUMER_PORT"`
		CELERYLOADER                  string `json:"CELERY_LOADER"`
		CELERYLOGFILE                 string `json:"CELERY_LOG_FILE"`
		CELERYLOGLEVEL                string `json:"CELERY_LOG_LEVEL"`
		CELERYLOGREDIRECT             string `json:"CELERY_LOG_REDIRECT"`
		CELERYLOGREDIRECTLEVEL        string `json:"CELERY_LOG_REDIRECT_LEVEL"`
		DJANGOLIVETESTSERVERADDRESS   string `json:"DJANGO_LIVE_TEST_SERVER_ADDRESS"`
		DJANGOPROJECTDIR              string `json:"DJANGO_PROJECT_DIR"`
		DJANGOSETTINGSMODULE          string `json:"DJANGO_SETTINGS_MODULE"`
		HOME                          string `json:"HOME"`
		INVENTORYHOSTVARS             string `json:"INVENTORY_HOSTVARS"`
		INVENTORYID                   string `json:"INVENTORY_ID"`
		LANG                          string `json:"LANG"`
		PATH                          string `json:"PATH"`
		PROOTTMPDIR                   string `json:"PROOT_TMP_DIR"`
		PS1                           string `json:"PS1"`
		PWD                           string `json:"PWD"`
		PYTHONPATH                    string `json:"PYTHONPATH"`
		RESTAPITOKEN                  string `json:"REST_API_TOKEN"`
		RESTAPIURL                    string `json:"REST_API_URL"`
		SHLVL                         string `json:"SHLVL"`
		SUPERVISORENABLED             string `json:"SUPERVISOR_ENABLED"`
		SUPERVISORGROUPNAME           string `json:"SUPERVISOR_GROUP_NAME"`
		SUPERVISORPROCESSNAME         string `json:"SUPERVISOR_PROCESS_NAME"`
		SUPERVISORSERVERURL           string `json:"SUPERVISOR_SERVER_URL"`
		TZ                            string `json:"TZ"`
		USER                          string `json:"USER"`
		VIRTUALENV                    string `json:"VIRTUAL_ENV"`
		MPFORKLOGFILE                 string `json:"_MP_FORK_LOGFILE_"`
		MPFORKLOGFORMAT               string `json:"_MP_FORK_LOGFORMAT_"`
		MPFORKLOGLEVEL                string `json:"_MP_FORK_LOGLEVEL_"`
	} `json:"job_env"`
	JobExplanation string `json:"job_explanation"`
	JobType        string `json:"job_type"`
	LaunchType     string `json:"launch_type"`
	Limit          string `json:"limit"`
	Modified       string `json:"modified"`
	ModuleArgs     string `json:"module_args"`
	ModuleName     string `json:"module_name"`
	Name           string `json:"name"`
	Related        struct {
		ActivityStream string `json:"activity_stream"`
		Cancel         string `json:"cancel"`
		CreatedBy      string `json:"created_by"`
		Credential     string `json:"credential"`
		Events         string `json:"events"`
		Inventory      string `json:"inventory"`
		Relaunch       string `json:"relaunch"`
		Stdout         string `json:"stdout"`
	} `json:"related"`
	ResultStdout    string `json:"result_stdout"`
	ResultTraceback string `json:"result_traceback"`
	Started         string `json:"started"`
	Status          string `json:"status"`
	SummaryFields   struct {
		CreatedBy struct {
			FirstName string `json:"first_name"`
			ID        int    `json:"id"`
			LastName  string `json:"last_name"`
			Username  string `json:"username"`
		} `json:"created_by"`
		Credential struct {
			Cloud       bool   `json:"cloud"`
			Description string `json:"description"`
			ID          int    `json:"id"`
			Kind        string `json:"kind"`
			Name        string `json:"name"`
		} `json:"credential"`
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
	} `json:"summary_fields"`
	Type      string `json:"type"`
	URL       string `json:"url"`
	Verbosity int    `json:"verbosity"`
}
