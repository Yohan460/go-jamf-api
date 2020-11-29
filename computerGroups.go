package jamf

import "fmt"

const uriComputerGroups = "/JSSResource/computergroups"

type ComputerGroupsResponse struct {
	Count   int                         `xml:"size"`
	Results []ComputerGroupListResponse `xml:"computer_group"`
}

type ComputerGroupListResponse struct {
	ID      int    `xml:"id,omitempty"`
	Name    string `xml:"name,omitempty"`
	IsSmart bool   `xml:"is_smart,omitempty"`
}

type ComputerGroup struct {
	ID            int                      `xml:"id"`
	Name          string                   `xml:"name"`
	IsSmart       bool                     `xml:"is_smart"`
	Site          Site                     `xml:"site"`
	Criteria      []ComputerGroupCriterion `xml:"criteria>criterion"`
	CriteriaCount int                      `xml:"criteria>size"`
	Computers     []ComputerGroupResp      `xml:"computers>computer"`
	ComputerCount int                      `xml:"computers>size"`
}

type ComputerGroupRequest struct {
	Name      string                      `xml:"name"`
	IsSmart   bool                        `xml:"is_smart"`
	Site      Site                        `xml:"site"`
	Criteria  []ComputerGroupCriterion    `xml:"criteria>criterion"`
	Computers []ComputerGroupListResponse `xml:"computers>computer,omitempty"`
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

func (c *Client) GetComputerGroupByName(name string) (data *ComputerGroup, err error) {
	id, err := c.GetComputerGroupIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetComputerGroup(id)
}

func (c *Client) GetComputerGroup(id int) (data *ComputerGroup, err error) {
	var out *ComputerGroup
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, id)
	err = c.doRequest("GET", uri, nil, &out)

	return out, err
}

func (c *Client) GetComputerGroups() (data *ComputerGroupsResponse, err error) {
	out := &ComputerGroupsResponse{}
	err = c.doRequest("GET", uriComputerGroups, nil, &out)

	return out, err
}

func (c *Client) CreateComputerGroup(computerGroupRequest *ComputerGroupRequest) error {

	reqBody := &struct {
		*ComputerGroupRequest
		XMLName struct{} `xml:"computer_group"`
	}{
		ComputerGroupRequest: computerGroupRequest,
	}

	return c.doRequest("POST", fmt.Sprintf("%s/id/0", uriComputerGroups), reqBody, &struct{}{})
}

func (c *Client) UpdateComputerGroup(d *ComputerGroup) error {
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, d.ID)

	reqBody := &struct {
		*ComputerGroup
		XMLName struct{} `xml:"computer_group"`
	}{
		ComputerGroup: d,
	}

	return c.doRequest("PUT", uri, reqBody, &struct{}{})
}

func (c *Client) DeleteComputerGroup(id int) error {
	uri := fmt.Sprintf("%s/id/%v", uriComputerGroups, id)
	return c.doRequest("DELETE", uri, nil, &struct{}{})
}
