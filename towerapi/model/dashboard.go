package model

type Dashboard struct {
	Credentials struct {
		Total int    `json:"total"`
		URL   string `json:"url"`
	} `json:"credentials"`
	Groups struct {
		FailuresURL     string `json:"failures_url"`
		InventoryFailed int    `json:"inventory_failed"`
		JobFailed       int    `json:"job_failed"`
		Total           int    `json:"total"`
		URL             string `json:"url"`
	} `json:"groups"`
	Hosts struct {
		Failed      int    `json:"failed"`
		FailuresURL string `json:"failures_url"`
		Total       int    `json:"total"`
		URL         string `json:"url"`
	} `json:"hosts"`
	Inventories struct {
		InventoryFailed          int    `json:"inventory_failed"`
		JobFailed                int    `json:"job_failed"`
		Total                    int    `json:"total"`
		TotalWithInventorySource int    `json:"total_with_inventory_source"`
		URL                      string `json:"url"`
	} `json:"inventories"`
	InventorySources struct {
		Ec2 struct {
			Failed      int    `json:"failed"`
			FailuresURL string `json:"failures_url"`
			Label       string `json:"label"`
			Total       int    `json:"total"`
			URL         string `json:"url"`
		} `json:"ec2"`
		Rax struct {
			Failed      int    `json:"failed"`
			FailuresURL string `json:"failures_url"`
			Label       string `json:"label"`
			Total       int    `json:"total"`
			URL         string `json:"url"`
		} `json:"rax"`
	} `json:"inventory_sources"`
	JobTemplates struct {
		Total int    `json:"total"`
		URL   string `json:"url"`
	} `json:"job_templates"`
	Jobs struct {
		Failed     int    `json:"failed"`
		FailureURL string `json:"failure_url"`
		Total      int    `json:"total"`
		URL        string `json:"url"`
	} `json:"jobs"`
	Organizations struct {
		Total int    `json:"total"`
		URL   string `json:"url"`
	} `json:"organizations"`
	Projects struct {
		Failed      int    `json:"failed"`
		FailuresURL string `json:"failures_url"`
		Total       int    `json:"total"`
		URL         string `json:"url"`
	} `json:"projects"`
	Related struct {
		JobsGraph string `json:"jobs_graph"`
	} `json:"related"`
	ScmTypes struct {
		Git struct {
			Failed      int    `json:"failed"`
			FailuresURL string `json:"failures_url"`
			Label       string `json:"label"`
			Total       int    `json:"total"`
			URL         string `json:"url"`
		} `json:"git"`
		Hg struct {
			Failed      int    `json:"failed"`
			FailuresURL string `json:"failures_url"`
			Label       string `json:"label"`
			Total       int    `json:"total"`
			URL         string `json:"url"`
		} `json:"hg"`
		Svn struct {
			Failed      int    `json:"failed"`
			FailuresURL string `json:"failures_url"`
			Label       string `json:"label"`
			Total       int    `json:"total"`
			URL         string `json:"url"`
		} `json:"svn"`
	} `json:"scm_types"`
	Teams struct {
		Total int    `json:"total"`
		URL   string `json:"url"`
	} `json:"teams"`
	Users struct {
		Total int    `json:"total"`
		URL   string `json:"url"`
	} `json:"users"`
}
