package jamf

import (
	"fmt"

	"github.com/bradfitz/slice"
)

const uriComputerGroups = "/JSSResource/computergroups"

type ComputerGroupsResponse struct {
	Size    int                         `xml:"size"`
	Results []ComputerGroupListResponse `xml:"computer_group"`
}

type ComputerGroupListResponse struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	IsSmart bool   `xml:"is_smart,omitempty"`
}

type ComputerGroup struct {
	ID           int                          `xml:"id"`
	Name         string                       `xml:"name"`
	IsSmart      bool                         `xml:"is_smart"`
	Site         Site                         `xml:"site"`
	Criteria     []ComputerGroupCriterion     `xml:"criteria>criterion"`
	CriteriaSize int                          `xml:"criteria>size"`
	Computers    []ComputerGroupComputerEntry `xml:"computers>computer"`
	ComputerSize int                          `xml:"computers>size"`
}

type ComputerGroupRequest struct {
	Name      string                       `xml:"name"`
	IsSmart   bool                         `xml:"is_smart"`
	Site      Site                         `xml:"site"`
	Criteria  []ComputerGroupCriterion     `xml:"criteria>criterion"`
	Computers []ComputerGroupComputerEntry `xml:"computers>computer,omitempty"`
}

type ComputerGroupComputerEntry struct {
	ID           int    `json:"id,omitempty" xml:"id,omitempty"`
	Name         string `json:"name,omitempty" xml:"name,omitempty"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number,omitempty"`
}

type ComputerGroupCriterion struct {
	Name         string             `xml:"name"`
	Priority     int                `xml:"priority"`
	AndOr        ComputerGroupAndOr `xml:"and_or"`
	SearchType   string             `xml:"search_type"`
	SearchValue  string             `xml:"value"`
	OpeningParen bool               `xml:"opening_paren"`
	ClosingParen bool               `xml:"closing_paren"`
}

type ComputerGroupAndOr string

const (
	And ComputerGroupAndOr = "and"
	Or  ComputerGroupAndOr = "or"
)

func (c *Client) GetComputerGroupIdByName(name string) (int, error) {
	var id int
	d, err := c.GetComputerGroups()
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

func (c *Client) GetComputerGroupByName(name string) (*ComputerGroup, error) {
	id, err := c.GetComputerGroupIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetComputerGroup(id)
}

func (c *Client) GetComputerGroup(id int) (*ComputerGroup, error) {
	var out *ComputerGroup
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, id)
	err := c.doRequest("GET", uri, nil, &out)

	return out, err
}

func (c *Client) GetComputerGroups() (*ComputerGroupsResponse, error) {
	out := &ComputerGroupsResponse{}
	err := c.doRequest("GET", uriComputerGroups, nil, out)

	return out, err
}

func (c *Client) CreateComputerGroup(d *ComputerGroupRequest) (int, error) {

	group := &ComputerGroup{}
	d.Criteria = validateComputerGroupCriteriaOrder(d.Criteria)
	reqBody := &struct {
		*ComputerGroupRequest
		XMLName struct{} `xml:"computer_group"`
	}{
		ComputerGroupRequest: d,
	}

	err := c.doRequest("POST", fmt.Sprintf("%s/id/0", uriComputerGroups), reqBody, group)
	return group.ID, err
}

func (c *Client) UpdateComputerGroup(d *ComputerGroup) (int, error) {

	group := &ComputerGroup{}
	d.Criteria = validateComputerGroupCriteriaOrder(d.Criteria)
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, d.ID)
	reqBody := &struct {
		*ComputerGroup
		XMLName struct{} `xml:"computer_group"`
	}{
		ComputerGroup: d,
	}

	err := c.doRequest("PUT", uri, reqBody, group)

	return group.ID, err
}

func (c *Client) DeleteComputerGroup(id int) (int, error) {

	group := &ComputerGroup{}
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, id)
	err := c.doRequest("DELETE", uri, nil, group)

	return group.ID, err
}

func validateComputerGroupCriteriaOrder(criteria []ComputerGroupCriterion) []ComputerGroupCriterion {
	slice.Sort(criteria[:], func(i, j int) bool {
		return criteria[i].Priority < criteria[j].Priority
	})
	return criteria
}
