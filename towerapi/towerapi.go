package towerapi

import (
	"fmt"
	"github.com/dghubble/sling"
	s "github.com/mpeter/go-towerapi/towerapi/services"
	"net/http"
	"net/url"
)

const (
	libraryVersion  = "0.1.0"
	userAgent       = "go-towerapi/" + libraryVersion
	mediaType       = "application/json"
	defaultEndpoint = "http://localhost/api/v1/"
	defaultUsername = "admin"
	defaultPassword = "admin"
)

// Client manages communication with Ansible Tower V1 API.
type Client struct {
	// HTTP client used to communicate with the Ansible Tower API.
	sling *sling.Sling

	// Username for initial authentication
	Username string

	// Password for initial authentication
	Password string

	// Base URL for API requests.
	Endpoint *url.URL

	// AuthToken obtained from calling Ansible Tower API /authtoken/ endpoint
	Token AuthToken

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	ActivityStream        *s.ActivityStreamService
	AdHocCommands         *s.AdHocCommandsService
	AuthToken             *s.AuthTokenService
	Config                *s.ConfigService
	Credentials           *s.CredentialsService
	Dashboard             *s.DashboardService
	Groups                *s.GroupsService
	Hosts                 *s.HostsService
	Inventories           *s.InventoriesService
	InventoryScripts      *s.InventoryScriptsService
	InventorySources      *s.InventorySourcesService
	JobEvents             *s.JobEventsService
	JobTemplates          *s.JobTemplatesService
	Jobs                  *s.JobsService
	Labels                *s.LabelsService
	Me                    *s.MeService
	NotificationTemplates *s.NotificationTemplatesService
	Notifications         *s.NotificationsService
	Organizations         *s.OrganizationsService
	Ping                  *s.PingService
	Projects              *s.ProjectsService
	Roles                 *s.RolesService
	Schedules             *s.SchedulesService
	SystemJobTemplates    *s.SystemJobTemplatesService
	SystemJobs            *s.SystemJobsService
	Teams                 *s.TeamsService
	UnifiedJobTemplates   *s.UnifiedJobTemplatesService
	UnifiedJobs           *s.UnifiedJobsService
	Users                 *s.UsersService

	// Optional function called after every successful request made to the Tower APIs
	onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
type RequestCompletionCallback func(*http.Request, *http.Response)

// Response is a Ansible Tower response. This wraps the standard http.Response returned from Ansible Tower.
type Response struct {
	*http.Response
}

// An ErrorResponse reports the error caused by an API request
type ErrorResponse struct {
	// HTTP response that caused this error
	Response *http.Response
	All      []string `json:"__all__"`
}

func NewClient(httpClient *http.Client, endpoint string, username string, password string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	ep, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	base := sling.New().Client(httpClient).Base(endpoint)

	body := &AuthTokenCreateRequest{
		Username: username,
		Password: password,
	}
	authToken := new(AuthToken)
	errorResponse := new(ErrorResponse)
	_, err = base.Post("authtoken/").BodyJSON(body).ReceiveSuccess(authToken, errorResponse)

	token := fmt.Sprintf("Token %s", authToken.Token)
	base.Set("User-Agent", userAgent)
	base.Set("Authorization", token)

	return &Client{
		sling:     base,
		Endpoint:  ep,
		Username:  username,
		Password:  password,
		Token:     token,
		UserAgent: userAgent,

		ActivityStream:   NewActivityStreamService(base.New()),
		AdHocCommands:    NewAdHocCommandsService(base.New()),
		AuthToken:        NewAuthTokenService(base.New()),
		Config:           NewConfigService(base.New()),
		Credentials:      NewCredentialsService(base.New()),
		Dashboard:        NewDashboardService(base.New()),
		Groups:           NewGroupsService(base.New()),
		Hosts:            NewHostsService(base.New()),
		Inventories:      NewInventoriesService(base.New()),
		InventoryScripts: NewInventoryScriptsService(base.New()),
		InventorySources: NewInventorySourcesService(base.New()),
		JobEvents:        NewJobEventsService(base.New()),
		JobTemplates:     NewJobTemplatesService(base.New()),
		Jobs:             NewJobsService(base.New()),
		Labels:           NewLablesService(base.New()),
		Me:               NewMeService(base.New()),
		NotificationTemplates: NewNotificationTemplatesService(base.New()),
		Notifications:         NewNotificationsService(base.New()),
		Organizations:         NewOrganizationsService(base.New()),
		Ping:                  NewPingService(base.New()),
		Projects:              NewProjectsService(base.New()),
		Roles:                 NewRolesService(base.New()),
		Schedules:             NewSchedulesService(base.New()),
		SystemJobTemplates:    NewSystemJobTemplatesService(base.New()),
		SystemJobs:            NewSystemJobsService(base.New()),
		Teams:                 NewTeamsService(base.New()),
		UnifiedJobTemplates:   NewUnifiedJobTemplatesService(base.New()),
		UnifiedJobs:           NewUnifiedJobsService(base.New()),
		Users:                 NewUsersService(base.New()),
	}

	return c, nil
}

func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {

}

func (c *Client) List(root interface{}, opt *ListOptions, path string) (*Response, error) {

	path, err := addOptions(path, opt)
	if err != nil {
		return nil, err
	}

	req, err := c.NewRequest("GET", path, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.Do(req, root)

	if err != nil {
		return resp, err
	}
	return resp, err
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v",
		r.Response.Request.Method, r.Response.Request.URL, r.Response.StatusCode, r.Message)
}
