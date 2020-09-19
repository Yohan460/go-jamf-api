package jamf

const uriDepartments = "/v1/departments"

type Department struct {
	TotalCount *int               `json:"totalCount,omitempty"`
	Results    []DepartmentResult `json:"results,omitempty"`
}

type DepartmentResult struct {
	Id   *string `json:"id,omitempty"` // The response type to be returned is a string
	Name *string `json:"name,omitempty"`
}

func (c *Client) GetDepartments() (data Department, err error) {
	var d Department
	err = c.doRequest("GET", uriDepartments, nil, &d)
	data = d
	return
}
