package towerapi

type Notifications struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []Notification `json:"results"`
}

type Notification struct {
	Created              string `json:"created"`
	Error                string `json:"error"`
	ID                   int    `json:"id"`
	Modified             string `json:"modified"`
	NotificationTemplate int    `json:"notification_template"`
	NotificationType     string `json:"notification_type"`
	NotificationsSent    int    `json:"notifications_sent"`
	Recipients           string `json:"recipients"`
	Related              struct {
		NotificationTemplate string `json:"notification_template"`
	} `json:"related"`
	Status        string `json:"status"`
	Subject       string `json:"subject"`
	SummaryFields struct {
		NotificationTemplate struct {
			Description string `json:"description"`
			ID          int    `json:"id"`
			Name        string `json:"name"`
		} `json:"notification_template"`
	} `json:"summary_fields"`
	Type string `json:"type"`
	URL  string `json:"url"`
}
