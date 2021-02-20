package jamf

import (
	"fmt"
)

const uriComputerConfigurationProfiles = "/JSSResource/osxconfigurationprofiles"

type ComputerConfigurationProfile struct {
	General     ComputerConfigurationProfileGeneral     `xml:"general"`
	Scope       ComputerConfigurationProfileScope       `xml:"scope"`
	SelfService ComputerConfigurationProfileSelfService `xml:"self_service,omitempty"`
}

type ComputerConfigurationProfileGeneral struct {
	ID                 int             `xml:"id,omitempty"`
	Name               string          `xml:"name"`
	Description        string          `xml:"description"`
	Site               Site            `xml:"site"`
	Category           GeneralCategory `xml:"category,omitempty"`
	DistributionMethod string          `xml:"distribution_method"`
	UserRemovable      bool            `xml:"user_removable"`
	Level              string          `xml:"level"`
	UUID               string          `xml:"uuid"`
	RedeployOnUpdate   string          `xml:"redeploy_on_update"`
	Payload            string          `xml:"payloads"`
}

type ComputerConfigurationProfileScope struct {
	AllComputers   bool                                         `xml:"all_computers"`
	AllUsers       bool                                         `xml:"all_jss_users"`
	Computers      []ComputerScope                              `xml:"computers>computer,omitempty"`
	Buildings      []BuildingScope                              `xml:"buildings>building,omitempty"`
	Departments    []DepartmentScope                            `xml:"departments>department,omitempty"`
	ComputerGroups []ComputerGroupListResponse                  `xml:"computer_groups>computer_group,omitempty"`
	JamfUsers      []JamfUserScope                              `xml:"jss_users>jss_user,omitempty"`
	JamfUserGroups []UserGroupScope                             `xml:"jss_user_groups>jss_user_group,omitempty"`
	Limitiations   ComputerConfigurationProfileScopeLimitations `xml:"limitations"`
	Exclusions     ComputerConfigurationProfileScopeExclusions  `xml:"exclusions"`
}

type ComputerConfigurationProfileScopeLimitations struct {
	Users           []UserScope           `xml:"users>user,omitempty"`
	UserGroups      []UserGroupScope      `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []NetworkSegmentScope `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []IBeaconScope        `xml:"ibeacons>ibeacon,omitempty"`
}
type ComputerConfigurationProfileScopeExclusions struct {
	Computers       []ComputerScope             `xml:"computers>computer,omitempty"`
	Buildings       []BuildingScope             `xml:"buildings>building,omitempty"`
	Departments     []DepartmentScope           `xml:"departments>department,omitempty"`
	ComputerGroups  []ComputerGroupListResponse `xml:"computer_groups>computer_group,omitempty"`
	Users           []UserScope                 `xml:"users>user,omitempty"`
	UserGroups      []UserGroupScope            `xml:"user_groups>user_group,omitempty"`
	NetworkSegments []NetworkSegmentScope       `xml:"network_segments>network_segment,omitempty"`
	IBeacons        []IBeaconScope              `xml:"ibeacons>ibeacon,omitempty"`
	JamfUsers       []JamfUserScope             `xml:"jss_users>jss_user,omitempty"`
	JamfUserGroups  []UserGroupScope            `xml:"jss_user_groups>jss_user_group,omitempty"`
}

type ComputerConfigurationProfileSelfService struct {
	SelfServiceDisplayName      string                `xml:"self_service_display_name,omitempty"`
	InstallButtonText           string                `xml:"install_button_text,omitempty"`
	SelfServiceDescription      string                `xml:"self_service_description,omitempty"`
	ForceUsersToViewDescription bool                  `xml:"force_users_to_view_description,omitempty"`
	RemovalDisallowed           string                `xml:"security>removal_disallowed,omitempty"`
	SelfServiceIcon             SelfServiceIcon       `xml:"self_service_icon,omitempty,omitempty"`
	FeatureOnMainPage           bool                  `xml:"feature_on_main_page,omitempty"`
	SelfServiceCategories       []SelfServiceCategory `xml:"self_service_categories>category,omitempty,omitempty"`
}

type ComputerConfigurationProfileListResponse struct {
	Size    int                                    `xml:"size"`
	Results []ComputerConfigurationProfileListItem `xml:"os_x_configuration_profile"`
}

type ComputerConfigurationProfileListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func (c *Client) GetComputerConfigurationProfileIdByName(name string) (int, error) {
	var id int
	d, err := c.GetComputerConfigurationProfiles()
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

func (c *Client) GetComputerConfigurationProfileByName(name string) (*ComputerConfigurationProfile, error) {
	id, err := c.GetComputerConfigurationProfileIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetComputerConfigurationProfile(id)
}

func (c *Client) GetComputerConfigurationProfile(id int) (*ComputerConfigurationProfile, error) {
	var out *ComputerConfigurationProfile
	uri := fmt.Sprintf("%s/id/%v", uriComputerConfigurationProfiles, id)
	err := c.doRequest("GET", uri, nil, &out)

	return out, err
}

func (c *Client) GetComputerConfigurationProfiles() (*ComputerConfigurationProfileListResponse, error) {
	out := &ComputerConfigurationProfileListResponse{}
	err := c.doRequest("GET", uriComputerConfigurationProfiles, nil, out)

	return out, err
}

func (c *Client) CreateComputerConfigurationProfile(d *ComputerConfigurationProfile) (int, error) {

	// Setting usual defaults
	if d.General.DistributionMethod == "" {
		d.General.DistributionMethod = "Install Automatically"
	}
	if d.General.Level == "" {
		d.General.Level = "computer"
	}

	res := &ComputerConfigurationProfileListItem{}
	reqBody := &struct {
		*ComputerConfigurationProfile
		XMLName struct{} `xml:"os_x_configuration_profile"`
	}{
		ComputerConfigurationProfile: d,
	}

	err := c.doRequest("POST", fmt.Sprintf("%s/id/0", uriComputerConfigurationProfiles), reqBody, res)
	return res.ID, err
}

func (c *Client) UpdateComputerConfigurationProfile(d *ComputerConfigurationProfile) (int, error) {

	res := &ComputerConfigurationProfile{}
	uri := fmt.Sprintf("%s/id/%v", uriComputerConfigurationProfiles, d.General.ID)
	reqBody := &struct {
		*ComputerConfigurationProfile
		XMLName struct{} `xml:"os_x_configuration_profile"`
	}{
		ComputerConfigurationProfile: d,
	}

	err := c.doRequest("PUT", uri, reqBody, res)

	return res.General.ID, err
}

func (c *Client) DeleteComputerConfigurationProfile(id int) (int, error) {

	profile := &ComputerConfigurationProfile{}
	uri := fmt.Sprintf("%s/id/%v", uriComputerConfigurationProfiles, id)
	err := c.doRequest("DELETE", uri, nil, profile)

	return profile.General.ID, err
}
