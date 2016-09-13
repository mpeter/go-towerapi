package towerapi

type Jobs struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []Job  `json:"results"`
}

type Job struct {
	AskCredentialOnLaunch bool    `json:"ask_credential_on_launch"`
	AskInventoryOnLaunch  bool    `json:"ask_inventory_on_launch"`
	AskJobTypeOnLaunch    bool    `json:"ask_job_type_on_launch"`
	AskLimitOnLaunch      bool    `json:"ask_limit_on_launch"`
	AskSkipTagsOnLaunch   bool    `json:"ask_skip_tags_on_launch"`
	AskTagsOnLaunch       bool    `json:"ask_tags_on_launch"`
	AskVariablesOnLaunch  bool    `json:"ask_variables_on_launch"`
	CloudCredential       string  `json:"cloud_credential"`
	Created               string  `json:"created"`
	Credential            int     `json:"credential"`
	Description           string  `json:"description"`
	Elapsed               float64 `json:"elapsed"`
	ExtraVars             string  `json:"extra_vars"`
	Failed                bool    `json:"failed"`
	Finished              string  `json:"finished"`
	ForceHandlers         bool    `json:"force_handlers"`
	Forks                 int     `json:"forks"`
	ID                    int     `json:"id"`
	Inventory             int     `json:"inventory"`
	JobArgs               string  `json:"job_args"`
	JobCwd                string  `json:"job_cwd"`
	JobEnv                struct {
		ANSIBLECALLBACKPLUGINS        string `json:"ANSIBLE_CALLBACK_PLUGINS"`
		ANSIBLEFORCECOLOR             string `json:"ANSIBLE_FORCE_COLOR"`
		ANSIBLEHOSTKEYCHECKING        string `json:"ANSIBLE_HOST_KEY_CHECKING"`
		ANSIBLEPARAMIKORECORDHOSTKEYS string `json:"ANSIBLE_PARAMIKO_RECORD_HOST_KEYS"`
		ANSIBLESSHCONTROLPATH         string `json:"ANSIBLE_SSH_CONTROL_PATH"`
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
		JOBID                         string `json:"JOB_ID"`
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
	JobExplanation         string        `json:"job_explanation"`
	JobTags                string        `json:"job_tags"`
	JobTemplate            int           `json:"job_template"`
	JobType                string        `json:"job_type"`
	LaunchType             string        `json:"launch_type"`
	Limit                  string        `json:"limit"`
	Modified               string        `json:"modified"`
	Name                   string        `json:"name"`
	NetworkCredential      string        `json:"network_credential"`
	PasswordsNeededToStart []interface{} `json:"passwords_needed_to_start"`
	Playbook               string        `json:"playbook"`
	Project                int           `json:"project"`
	Related                struct {
		ActivityStream     string `json:"activity_stream"`
		Cancel             string `json:"cancel"`
		CreatedBy          string `json:"created_by"`
		Credential         string `json:"credential"`
		Inventory          string `json:"inventory"`
		JobEvents          string `json:"job_events"`
		JobHostSummaries   string `json:"job_host_summaries"`
		JobPlays           string `json:"job_plays"`
		JobTasks           string `json:"job_tasks"`
		JobTemplate        string `json:"job_template"`
		Labels             string `json:"labels"`
		Notifications      string `json:"notifications"`
		Project            string `json:"project"`
		Relaunch           string `json:"relaunch"`
		Start              string `json:"start"`
		Stdout             string `json:"stdout"`
		UnifiedJobTemplate string `json:"unified_job_template"`
	} `json:"related"`
	ResultStdout    string `json:"result_stdout"`
	ResultTraceback string `json:"result_traceback"`
	SkipTags        string `json:"skip_tags"`
	StartAtTask     string `json:"start_at_task"`
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
		JobTemplate struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"job_template"`
		Labels struct {
			Count   int           `json:"count"`
			Results []interface{} `json:"results"`
		} `json:"labels"`
		Project struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
			Status      string `json:"status"`
		} `json:"project"`
		UnifiedJobTemplate struct {
			Description    string `json:"description"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			UnifiedJobType string `json:"unified_job_type"`
		} `json:"unified_job_template"`
	} `json:"summary_fields"`
	Type               string `json:"type"`
	UnifiedJobTemplate int    `json:"unified_job_template"`
	URL                string `json:"url"`
	Verbosity          int    `json:"verbosity"`
}
