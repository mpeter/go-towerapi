package towerapi

type UnifiedJobs struct {
	Count    int          `json:"count"`
	Next     string       `json:"next"`
	Previous string       `json:"previous"`
	Results  []UnifiedJob `json:"results"`
}

type UnifiedJob struct {
	Created     string  `json:"created"`
	Description string  `json:"description"`
	Elapsed     float64 `json:"elapsed"`
	ExtraVars   string  `json:"extra_vars"`
	Failed      bool    `json:"failed"`
	Finished    string  `json:"finished"`
	ID          int     `json:"id"`
	JobArgs     string  `json:"job_args"`
	JobCwd      string  `json:"job_cwd"`
	JobEnv      struct {
		ANSIBLEFORCECOLOR             string `json:"ANSIBLE_FORCE_COLOR"`
		ANSIBLEHOSTKEYCHECKING        string `json:"ANSIBLE_HOST_KEY_CHECKING"`
		ANSIBLEPARAMIKORECORDHOSTKEYS string `json:"ANSIBLE_PARAMIKO_RECORD_HOST_KEYS"`
		ANSIBLEUSEVENV                string `json:"ANSIBLE_USE_VENV"`
		ANSIBLEVENVPATH               string `json:"ANSIBLE_VENV_PATH"`
		CELERYLOADER                  string `json:"CELERY_LOADER"`
		CELERYLOGFILE                 string `json:"CELERY_LOG_FILE"`
		CELERYLOGLEVEL                string `json:"CELERY_LOG_LEVEL"`
		CELERYLOGREDIRECT             string `json:"CELERY_LOG_REDIRECT"`
		CELERYLOGREDIRECTLEVEL        string `json:"CELERY_LOG_REDIRECT_LEVEL"`
		DJANGOLIVETESTSERVERADDRESS   string `json:"DJANGO_LIVE_TEST_SERVER_ADDRESS"`
		DJANGOPROJECTDIR              string `json:"DJANGO_PROJECT_DIR"`
		DJANGOSETTINGSMODULE          string `json:"DJANGO_SETTINGS_MODULE"`
		HOME                          string `json:"HOME"`
		LANG                          string `json:"LANG"`
		PATH                          string `json:"PATH"`
		PS1                           string `json:"PS1"`
		PWD                           string `json:"PWD"`
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
	Modified       string `json:"modified"`
	Name           string `json:"name"`
	Related        struct {
		Cancel             string `json:"cancel"`
		Notifications      string `json:"notifications"`
		SystemJobTemplate  string `json:"system_job_template"`
		UnifiedJobTemplate string `json:"unified_job_template"`
	} `json:"related"`
	ResultStdout    string `json:"result_stdout"`
	ResultTraceback string `json:"result_traceback"`
	Started         string `json:"started"`
	Status          string `json:"status"`
	SummaryFields   struct {
		UnifiedJobTemplate struct {
			Description    string `json:"description"`
			ID             int    `json:"id"`
			Name           string `json:"name"`
			UnifiedJobType string `json:"unified_job_type"`
		} `json:"unified_job_template"`
	} `json:"summary_fields"`
	SystemJobTemplate  int    `json:"system_job_template"`
	Type               string `json:"type"`
	UnifiedJobTemplate int    `json:"unified_job_template"`
	URL                string `json:"url"`
}
