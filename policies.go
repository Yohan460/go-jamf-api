package jamf

import (
	"fmt"
)

const uriPolicys = "/JSSResource/policies"

type Policy struct {
	General              PolicyGeneral              `xml:"general"`
	Scope                PolicyScope                `xml:"scope,omitempty"`
	SelfService          PolicySelfService          `xml:"self_service"`
	PackageConfiguration PolicyPackageConfiguration `xml:"package_configuration,omitempty"`
	Reboot               PolicyReboot               `xml:"reboot"`
	Maintenance          PolicyMaintenance          `xml:"maintenance"`
	FilesAndProcesses    PolicyFilesAndProcesses    `xml:"files_processes"`
	UserInteraction      PolicyUserInteraction      `xml:"user_interaction"`
}

type PolicyGeneral struct {
	ID                         int                                  `xml:"id"`
	Name                       string                               `xml:"name"`
	Enabled                    bool                                 `xml:"enabled"`
	Trigger                    string                               `xml:"trigger"`
	TriggerCheckin             bool                                 `xml:"trigger_checkin"`
	TriggerEnrollmentComplete  bool                                 `xml:"trigger_enrollment_complete"`
	TriggerLogin               bool                                 `xml:"trigger_login"`
	TriggerLogout              bool                                 `xml:"trigger_logout"`
	TriggerNetworkStateChanged bool                                 `xml:"trigger_network_state_changed"`
	TriggerStartup             bool                                 `xml:"trigger_startup"`
	TriggerOther               string                               `xml:"trigger_other"`
	Frequency                  string                               `xml:"frequency"`
	RetryEvent                 string                               `xml:"retry_event"`
	RetryAttempts              int                                  `xml:"retry_attempts"`
	NotifyOnEachFailedRetry    bool                                 `xml:"notify_on_each_failed_retry"`
	LocationUserOnly           bool                                 `xml:"location_user_only"`
	TargetDrive                string                               `xml:"target_drive"`
	Offline                    bool                                 `xml:"offline"`
	Category                   PolicyGeneralCategory                `xml:"category"`
	DateTimeLimitations        PolicyGeneralDateTimeLimitations     `xml:"date_time_limitations"`
	NetworkLimitations         PolicyGeneralNetworkLimitations      `xml:"network_limitations"`
	OverrideDefaultSettings    PolicyGeneralOverrideDefaultSettings `xml:"override_default_settings"`
	NetworkRequirements        string                               `xml:"network_requirements"`
	Site                       Site                                 `xml:"site"`
}

type PolicyGeneralCategory struct {
	ID   string `xml:"id"`
	Name string `xml:"name"`
}

// TODO Get types
type PolicyGeneralDateTimeLimitations struct {
	ActivationDate      string `xml:"activation_date"`
	ActivationDateEpoch string `xml:"activation_date_epoch"`
	ActivationDateUtc   string `xml:"activation_date_utc"`
	ExpirationDate      string `xml:"expiration_date"`
	ExpirationDateEpoch string `xml:"expiration_date_epoch"`
	ExpirationDateUtc   string `xml:"expiration_date_utc"`
	NoExecuteOn         string `xml:"no_execute_on"`
	NoExecuteStart      string `xml:"no_execute_start"`
	NoExecuteEnd        string `xml:"no_execute_end"`
}

// TODO Get types
type PolicyGeneralNetworkLimitations struct {
	MinimumNetworkConnection string `xml:"minimum_network_connection"`
	AnyIpAddress             string `xml:"any_ip_address"`
	NetworkSegments          string `xml:"network_segments"`
}

// TODO Get types
type PolicyGeneralOverrideDefaultSettings struct {
	TargetDrive       string `xml:"target_drive"`
	DistributionPoint string `xml:"distribution_point"`
	ForceAfpSmb       string `xml:"force_afp_smb"`
	Sus               string `xml:"sus"`
	NetbootServer     string `xml:"netboot_server"`
}

type PolicyScope struct {
	AllComputers   string                    `xml:"all_computers"`
	Computers      string                    `xml:"computers"`
	ComputerGroups PolicyScopeComputerGroups `xml:"computer_groups"`
	// Buildings      string                    `xml:"buildings"`
	// Departments    string                    `xml:"departments"`
	// LimitToUsers   struct {
	// 	UserGroups string `xml:"user_groups"`
	// } `xml:"limit_to_users"`
	Limitations PolicyScopeLimitations `xml:"limitations"`
	Exclusions  PolicyScopeExclusions  `xml:"exclusions"`
}

type PolicyScopeComputerGroups struct {
	ComputerGroups []ComputerGroupListResponse `xml:"computer_group"`
}

type PolicySelfService struct {
	UseForSelfService           string `xml:"use_for_self_service"`
	SelfServiceDisplayName      string `xml:"self_service_display_name"`
	InstallButtonText           string `xml:"install_button_text"`
	ReinstallButtonText         string `xml:"reinstall_button_text"`
	SelfServiceDescription      string `xml:"self_service_description"`
	ForceUsersToViewDescription string `xml:"force_users_to_view_description"`
	SelfServiceIcon             string `xml:"self_service_icon"`
	FeatureOnMainPage           string `xml:"feature_on_main_page"`
	SelfServiceCategories       string `xml:"self_service_categories"`
}

type PolicyScopeLimitations struct {
	Users           string `xml:"users"`
	UserGroups      string `xml:"user_groups"`
	NetworkSegments string `xml:"network_segments"`
	Ibeacons        string `xml:"ibeacons"`
}

type PolicyScopeExclusions struct {
	Computers       string `xml:"computers"`
	ComputerGroups  string `xml:"computer_groups"`
	Buildings       string `xml:"buildings"`
	Departments     string `xml:"departments"`
	Users           string `xml:"users"`
	UserGroups      string `xml:"user_groups"`
	NetworkSegments string `xml:"network_segments"`
	Ibeacons        string `xml:"ibeacons"`
}

type PolicyPackageConfiguration struct {
	Packages []PolicyPackageConfigurationPackage `xml:"packages>package,omitempty"`
}

type PolicyPackageConfigurationPackage struct {
	ID int `xml:"id,omitempty"`
}

type PolicyReboot struct {
	Message                     string `xml:"message"`
	StartupDisk                 string `xml:"startup_disk"`
	SpecifyStartup              string `xml:"specify_startup"`
	NoUserLoggedIn              string `xml:"no_user_logged_in"`
	UserLoggedIn                string `xml:"user_logged_in"`
	MinutesUntilReboot          int    `xml:"minutes_until_reboot"`
	StartRebootTimerImmediately bool   `xml:"start_reboot_timer_immediately"`
	FileVault2Reboot            bool   `xml:"file_vault_2_reboot"`
}

type PolicyMaintenance struct {
	Recon                    string `xml:"recon"`
	ResetName                string `xml:"reset_name"`
	InstallAllCachedPackages string `xml:"install_all_cached_packages"`
	Heal                     string `xml:"heal"`
	Prebindings              string `xml:"prebindings"`
	Permissions              string `xml:"permissions"`
	Byhost                   string `xml:"byhost"`
	SystemCache              string `xml:"system_cache"`
	UserCache                string `xml:"user_cache"`
	Verify                   string `xml:"verify"`
}

type PolicyFilesAndProcesses struct {
	SearchByPath         string `xml:"search_by_path"`
	DeleteFile           string `xml:"delete_file"`
	LocateFile           string `xml:"locate_file"`
	UpdateLocateDatabase string `xml:"update_locate_database"`
	SpotlightSearch      string `xml:"spotlight_search"`
	SearchForProcess     string `xml:"search_for_process"`
	KillProcess          string `xml:"kill_process"`
	RunCommand           string `xml:"run_command"`
}

type PolicyUserInteraction struct {
	MessageStart          string `xml:"message_start"`
	AllowUsersToDefer     string `xml:"allow_users_to_defer"`
	AllowDeferralUntilUtc string `xml:"allow_deferral_until_utc"`
	AllowDeferralMinutes  string `xml:"allow_deferral_minutes"`
	MessageFinish         string `xml:"message_finish"`
}

type PolicyListResponse struct {
	Size    int              `xml:"size"`
	Results []PolicyListItem `xml:"policy"`
}

type PolicyListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func (c *Client) GetPolicyIdByName(name string) (int, error) {
	var id int
	d, err := c.GetPolicies()
	if err != nil {
		return -1, err
	}

	for _, v := range d.Results {
		if v.Name == name {
			id = v.ID
			break
		}
	}
	return id, err
}

func (c *Client) GetPolicyByName(name string) (*Policy, error) {
	id, err := c.GetPolicyIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetPolicy(id)
}

func (c *Client) GetPolicy(id int) (*Policy, error) {
	var out *Policy
	uri := fmt.Sprintf("%s/id/%v", uriPolicys, id)
	err := c.doRequest("GET", uri, nil, &out)

	return out, err
}

func (c *Client) GetPolicies() (*PolicyListResponse, error) {
	out := &PolicyListResponse{}
	err := c.doRequest("GET", uriPolicys, nil, out)

	return out, err
}

func (c *Client) CreatePolicy(d *Policy) (int, error) {

	// Setting defaults, per jamf unwritten requirements :/
	if d.General.Category.ID == "" || d.General.Category.Name == "" {
		d.General.Category.ID = "-1"
		d.General.Category.Name = "No category assigned"
	}
	if d.Reboot.StartupDisk == "" {
		d.Reboot.StartupDisk = "Current Startup Disk"
	}
	if d.Reboot.NoUserLoggedIn == "" {
		d.Reboot.NoUserLoggedIn = "Do not restart"
	}
	if d.Reboot.UserLoggedIn == "" {
		d.Reboot.UserLoggedIn = "Do not restart"
	}
	if d.Reboot.MinutesUntilReboot == 0 {
		d.Reboot.MinutesUntilReboot = 5
	}

	res := &PolicyListItem{}
	reqBody := &struct {
		*Policy
		XMLName struct{} `xml:"policy"`
	}{
		Policy: d,
	}

	err := c.doRequest("POST", fmt.Sprintf("%s/id/0", uriPolicys), reqBody, res)
	return res.ID, err
}

func (c *Client) UpdatePolicy(d *Policy) (int, error) {

	res := &PolicyListItem{}
	uri := fmt.Sprintf("%s/id/%v", uriPolicys, d.General.ID)
	reqBody := &struct {
		*Policy
		XMLName struct{} `xml:"policy"`
	}{
		Policy: d,
	}

	err := c.doRequest("PUT", uri, reqBody, res)

	return res.ID, err
}

func (c *Client) DeletePolicy(id int) (int, error) {

	policy := &Policy{}
	uri := fmt.Sprintf("%s/id/%v", uriPolicys, id)
	err := c.doRequest("DELETE", uri, nil, policy)

	return policy.General.ID, err
}
