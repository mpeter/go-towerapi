package towerapi

type Ping struct {
	Ha        bool `json:"ha"`
	Instances struct {
		Primary     string        `json:"primary"`
		Secondaries []interface{} `json:"secondaries"`
	} `json:"instances"`
	Role    string `json:"role"`
	Version string `json:"version"`
}
