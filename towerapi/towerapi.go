package towerapi

import (
	"net/http"

	"os"

	"github.com/mpeter/go-towerapi/towerapi/authtoken"
	"github.com/mpeter/go-towerapi/towerapi/credentials"
	"github.com/mpeter/go-towerapi/towerapi/errors"
	"github.com/mpeter/go-towerapi/towerapi/groups"
	"github.com/mpeter/go-towerapi/towerapi/hosts"
	"github.com/mpeter/go-towerapi/towerapi/inventories"
	"github.com/mpeter/go-towerapi/towerapi/job_templates"
	"github.com/mpeter/go-towerapi/towerapi/organizations"
	"github.com/mpeter/go-towerapi/towerapi/projects"
	"github.com/mpeter/sling"
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

	// AuthToken obtained from calling Ansible Tower API /authtoken/ endpoint
	Token *authtoken.AuthToken

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	AuthToken     *authtoken.Service
	Credentials   *credentials.Service
	Groups        *groups.Service
	Hosts         *hosts.Service
	Inventories   *inventories.Service
	JobTemplates  *job_templates.Service
	Organizations *organizations.Service
	Projects      *projects.Service

	//ActivityStream        *ActivityStreamService
	//AdHocCommands         *AdHocCommandsService
	//Config                *ConfigService
	//Dashboard             *DashboardService
	//InventoryScripts      *InventoryScriptsService
	//InventorySources      *InventorySourcesService
	//JobEvents             *JobEventsService
	//Jobs                  *JobsService
	//Labels                *LabelsService
	//Me                    *MeService
	//NotificationTemplates *NotificationTemplatesService
	//Notifications         *NotificationsService
	//Ping                  *PingService
	//Projects              *ProjectsService
	//Roles                 *RolesService
	//Schedules             *SchedulesService
	//SystemJobTemplates    *SystemJobTemplatesService
	//SystemJobs            *SystemJobsService
	//Teams                 *TeamsService
	//UnifiedJobTemplates   *UnifiedJobTemplatesService
	//UnifiedJobs           *UnifiedJobsService
	//Users                 *UsersService

}

type Config struct {
	Endpoint   string
	Username   string
	Password   string
	AuthToken  *authtoken.AuthToken
	HttpClient *http.Client
}

func DefaultConfig() *Config {

	config := &Config{
		Endpoint:   "http://127.0.0.1/api/v1/",
		Username:   "admin",
		Password:   "",
		AuthToken:  nil,
		HttpClient: http.DefaultClient,
	}

	if e := os.Getenv("TOWER_ENDOINT"); e != "" {
		config.Endpoint = e
	}
	if u := os.Getenv("TOWER_USERNAME"); u != "" {
		config.Username = u
	}
	if p := os.Getenv("TOWER_PASSWORD"); p != "" {
		config.Password = p
	}

	return config
}

func (c *Config) LoadAndValidate() error {

	base := sling.New().Client(c.HttpClient).Base(c.Endpoint)
	body := &authtoken.CreateRequest{
		Username: c.Username,
		Password: c.Password,
	}
	token := new(authtoken.AuthToken)
	apierr := new(errors.APIError)
	_, err := base.Post("authtoken/").BodyJSON(body).Receive(token, apierr)
	if error := errors.BuildError(err, apierr); error != nil {
		return error
	}
	c.AuthToken = token
	return nil
}

func NewClient(c *Config) (*Client, error) {

	if c.HttpClient == nil {
		c.HttpClient = http.DefaultClient
	}

	base := sling.New().Client(c.HttpClient).Base(c.Endpoint)
	base.Set("Authorization", "Token "+c.AuthToken.Token)

	return &Client{
		sling: base,

		AuthToken:     authtoken.NewService(base.New()),
		Credentials:   credentials.NewService(base.New()),
		Groups:        groups.NewService(base.New()),
		Hosts:         hosts.NewService(base.New()),
		Inventories:   inventories.NewService(base.New()),
		JobTemplates:  job_templates.NewService(base.New()),
		Organizations: organizations.NewService(base.New()),
		Projects:      projects.NewService(base.New()),

		//ActivityStream:   NewActivityStreamService(base),
		//AdHocCommands:    NewAdHocCommandsService(base),
		//Config:           NewConfigService(base),
		//Dashboard:        NewDashboardService(base),
		//InventoryScripts: NewInventoryScriptsService(base),
		//InventorySources: NewInventorySourcesService(base),
		//JobEvents:        NewJobEventsService(base),
		//Jobs:             NewJobsService(base),
		//Labels:           NewLabelsService(base),
		//Me:               NewMeService(base),
		//NotificationTemplates: NewNotificationTemplatesService(base),
		//Notifications:         NewNotificationsService(base),
		//Ping:                  NewPingService(base),
		//Roles:                 NewRolesService(base),
		//Schedules:             NewSchedulesService(base),
		//SystemJobTemplates:    NewSystemJobTemplatesService(base),
		//SystemJobs:            NewSystemJobsService(base),
		//Teams:                 NewTeamsService(base),
		//UnifiedJobTemplates:   NewUnifiedJobTemplatesService(base),
		//UnifiedJobs:           NewUnifiedJobsService(base),
		//Users:                 NewUsersService(base),
	}, nil
}
