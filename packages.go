package jamf

import (
	"fmt"
)

const uriPackages = "/JSSResource/packages"

type Package struct {
	ID                 int    `xml:"id"`
	Name               string `xml:"name"`
	CategoryName       string `xml:"category"`
	Filename           string `xml:"filename"`
	Info               string `xml:"info"`
	Notes              string `xml:"notes"`
	RebootRequired     bool   `xml:"reboot_required"`
	Priority           int    `xml:"priority"`
	FillExistingUsers  bool   `xml:"fill_existing_users"`
	BootVolumeRequired bool   `xml:"boot_volume_required"`
	AllowUninstalled   bool   `xml:"allow_uninstalled"`
	OsRequirements     string `xml:"os_requirements"`
	RequiredProcessor  string `xml:"required_processor"`
	HashType           string `xml:"hash_type,omitempty"`
	HashValue          string `xml:"hash_value,omitempty"`
}
type PackageListResponse struct {
	Size    int               `xml:"size"`
	Results []PackageListItem `xml:"package"`
}

type PackageListItem struct {
	ID   int    `xml:"id"`
	Name string `xml:"name"`
}

func (c *Client) GetPackageIdByName(name string) (int, error) {
	var id int
	d, err := c.GetPackages()
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

func (c *Client) GetPackageByName(name string) (*Package, error) {
	id, err := c.GetPackageIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetPackage(id)
}

func (c *Client) GetPackage(id int) (*Package, error) {
	var out *Package
	uri := fmt.Sprintf("%s/id/%v", uriPackages, id)
	err := c.doRequest("GET", uri, nil, &out)

	return out, err
}

func (c *Client) GetPackages() (*PackageListResponse, error) {
	out := &PackageListResponse{}
	err := c.doRequest("GET", uriPackages, nil, out)

	return out, err
}

func (c *Client) CreatePackage(d *Package) (int, error) {

	// Forcing these to empty strings to not allow setting
	d.HashType = ""
	d.HashValue = ""

	// Setting defaults, per jamf unwritten requirements :/
	if d.Priority == 0 {
		d.Priority = 10
	}
	if d.RequiredProcessor == "" {
		d.RequiredProcessor = "None"
	}
	if d.Filename == "" {
		d.Filename = "ThisPackageDoesNotExist.pkg"
	}

	res := &Package{}
	reqBody := &struct {
		*Package
		XMLName struct{} `xml:"package"`
	}{
		Package: d,
	}

	err := c.doRequest("POST", fmt.Sprintf("%s/id/0", uriPackages), reqBody, res)
	return res.ID, err
}

func (c *Client) UpdatePackage(d *Package) (int, error) {

	// Forcing these to empty strings to not allow updates
	d.HashType = ""
	d.HashValue = ""

	res := &Package{}
	uri := fmt.Sprintf("%s/id/%v", uriPackages, d.ID)
	reqBody := &struct {
		*Package
		XMLName struct{} `xml:"package"`
	}{
		Package: d,
	}

	err := c.doRequest("PUT", uri, reqBody, res)

	return res.ID, err
}

func (c *Client) DeletePackage(id int) (int, error) {

	group := &Package{}
	uri := fmt.Sprintf("%s/id/%v", uriPackages, id)
	err := c.doRequest("DELETE", uri, nil, group)

	return group.ID, err
}
