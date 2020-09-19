package jamf

const uriDepartments = "/v1/departments"

type Department struct {
	TotalCount *int               `json:"totalCount,omitempty"`
	Results    []DepartmentResult `json:"results,omitempty"`
}

type DepartmentResult struct {
	Id   *string `json:"id,omitempty"` // The response type to be returned is a string
	Name *string `json:"name,omitempty"`
	Href *string `json:"href,omitempty"`
}

func (c *Client) GetDepartments() (data Department, err error) {
	var d Department
	err = c.doRequest("GET", uriDepartments, nil, &d)
	data = d
	return
}

func (c *Client) CreateDepartment(name *string) (*DepartmentResult, error) {
	in := struct {
		Name *string `json:"name"`
	}{
		Name: name,
	}

	var out *DepartmentResult

	if err := c.doRequest("POST", uriDepartments, in, &out); err != nil {
		return nil, err
	}
	return out, nil
}
