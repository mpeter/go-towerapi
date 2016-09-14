package towerapi

import (
	"fmt"
	"net/http"

	"github.com/mpeter/go-towerapi/towerapi/authtoken"
	"github.com/mpeter/go-towerapi/towerapi/groups"
	"github.com/mpeter/go-towerapi/towerapi/hosts"
	"github.com/mpeter/go-towerapi/towerapi/inventories"
	"github.com/mpeter/go-towerapi/towerapi/organizations"
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
	Hosts         *hosts.Service
	Inventories   *inventories.Service
	Organizations *organizations.Service
	Groups        *groups.Service

	//ActivityStream        *ActivityStreamService
	//AdHocCommands         *AdHocCommandsService
	//Config                *ConfigService
	//Credentials           *CredentialsService
	//Dashboard             *DashboardService
	//InventoryScripts      *InventoryScriptsService
	//InventorySources      *InventorySourcesService
	//JobEvents             *JobEventsService
	//JobTemplates          *JobTemplatesService
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

	// Optional function called after every successful request made to the Tower APIs
	//onRequestCompleted RequestCompletionCallback
}

// RequestCompletionCallback defines the type of the request callback function
//type RequestCompletionCallback func(*http.Request, *http.Response)

type ClientConfig struct {
	Endpoint string
	Username string
	Password string
}

func NewClient(httpClient *http.Client, c *ClientConfig) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	base := sling.New().Client(httpClient).Base(c.Endpoint)
	body := &authtoken.CreateRequest{
		Username: c.Username,
		Password: c.Password,
	}
	ats := authtoken.NewService(base)
	at, err := ats.Create(body)
	if err != nil {
		return nil, fmt.Errorf("ERROR: Error Creating Token: %+v", err.Error())
	}
	token := fmt.Sprintf("Token %s", at.Token)
	base.Set("Authorization", token)

	return &Client{
		sling:         base,
		Token:         at,
		UserAgent:     userAgent,
		AuthToken:     authtoken.NewService(base.New()),
		Hosts:         hosts.NewService(base.New()),
		Inventories:   inventories.NewService(base.New()),
		Organizations: organizations.NewService(base.New()),
		Groups:        groups.NewService(base.New()),

		//ActivityStream:   NewActivityStreamService(base),
		//AdHocCommands:    NewAdHocCommandsService(base),
		//Config:           NewConfigService(base),
		//Credentials:      NewCredentialsService(base),
		//Dashboard:        NewDashboardService(base),
		//InventoryScripts: NewInventoryScriptsService(base),
		//InventorySources: NewInventorySourcesService(base),
		//JobEvents:        NewJobEventsService(base),
		//JobTemplates:     NewJobTemplatesService(base),
		//Jobs:             NewJobsService(base),
		//Labels:           NewLabelsService(base),
		//Me:               NewMeService(base),
		//NotificationTemplates: NewNotificationTemplatesService(base),
		//Notifications:         NewNotificationsService(base),
		//Ping:                  NewPingService(base),
		//Projects:              NewProjectsService(base),
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
