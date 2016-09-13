package towerapi

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/dghubble/sling"
	"github.com/mpeter/go-towerapi/towerapi/authtoken"
	"github.com/mpeter/go-towerapi/towerapi/hosts"
	"github.com/mpeter/go-towerapi/towerapi/inventories"
	"github.com/mpeter/go-towerapi/towerapi/organizations"
	"log"
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
	Token *authtoken.AuthToken

	// User agent for client
	UserAgent string

	// Services used for communicating with the API
	//ActivityStream        *ActivityStreamService
	//AdHocCommands         *AdHocCommandsService
	AuthToken *authtoken.Service
	//Config                *ConfigService
	//Credentials           *CredentialsService
	//Dashboard             *DashboardService
	//Groups                *GroupsService
	Hosts       *hosts.Service
	Inventories *inventories.Service
	//InventoryScripts      *InventoryScriptsService
	//InventorySources      *InventorySourcesService
	//JobEvents             *JobEventsService
	//JobTemplates          *JobTemplatesService
	//Jobs                  *JobsService
	//Labels                *LabelsService
	//Me                    *MeService
	//NotificationTemplates *NotificationTemplatesService
	//Notifications         *NotificationsService
	Organizations *organizations.Service
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

func NewClient(httpClient *http.Client, endpoint string, username string, password string) (*Client, error) {

	if httpClient == nil {
		httpClient = http.DefaultClient
	}

	ep, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}
	base := sling.New().Client(httpClient).Base(endpoint)

	body := &authtoken.CreateRequest{
		Username: username,
		Password: password,
	}
	ats := authtoken.NewService(base.New())
	authToken, _, err := ats.Create(body)
	if err != nil {
		log.Printf("ERROR: authToken.Create: %+v", err.Error())
		return nil, err
	}
	token := fmt.Sprintf("Token %s", authToken.Token)
	log.Printf("Token: %v", token)
	//base.Set("User-Agent", userAgent)
	base.Set("Authorization", token)

	return &Client{
		sling:     base,
		Endpoint:  ep,
		Username:  username,
		Password:  password,
		Token:     authToken,
		UserAgent: userAgent,

		//ActivityStream:   NewActivityStreamService(base.New()),
		//AdHocCommands:    NewAdHocCommandsService(base.New()),
		AuthToken: ats,
		//Config:           NewConfigService(base.New()),
		//Credentials:      NewCredentialsService(base.New()),
		//Dashboard:        NewDashboardService(base.New()),
		//Groups:           NewGroupsService(base.New()),
		Hosts:       hosts.NewService(base.New()),
		Inventories: inventories.NewService(base.New()),
		//InventoryScripts: NewInventoryScriptsService(base.New()),
		//InventorySources: NewInventorySourcesService(base.New()),
		//JobEvents:        NewJobEventsService(base.New()),
		//JobTemplates:     NewJobTemplatesService(base.New()),
		//Jobs:             NewJobsService(base.New()),
		//Labels:           NewLabelsService(base.New()),
		//Me:               NewMeService(base.New()),
		//NotificationTemplates: NewNotificationTemplatesService(base.New()),
		//Notifications:         NewNotificationsService(base.New()),
		Organizations: organizations.NewService(base.New()),
		//Ping:                  NewPingService(base.New()),
		//Projects:              NewProjectsService(base.New()),
		//Roles:                 NewRolesService(base.New()),
		//Schedules:             NewSchedulesService(base.New()),
		//SystemJobTemplates:    NewSystemJobTemplatesService(base.New()),
		//SystemJobs:            NewSystemJobsService(base.New()),
		//Teams:                 NewTeamsService(base.New()),
		//UnifiedJobTemplates:   NewUnifiedJobTemplatesService(base.New()),
		//UnifiedJobs:           NewUnifiedJobsService(base.New()),
		//Users:                 NewUsersService(base.New()),
	}, nil
}
