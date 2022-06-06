package jamf

import "fmt"

const uriDepartments = "/uapi/v1/departments"

type ResponseDepartments struct {
	TotalCount *int         `json:"totalCount,omitempty"`
	Results    []Department `json:"results,omitempty"`
}

type Department struct {
	Id   *string `json:"id,omitempty"` // The response type to be returned is a string
	Name *string `json:"name,omitempty"`
	Href *string `json:"href,omitempty"`
}

type DepartmentScope struct {
	ID   int    `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

func (c *Client) GetDepartmentIdByName(name string) (string, error) {
	var id string
	d, err := c.GetDepartments()
	if err != nil {
		return "", err
	}

	for _, v := range d.Results {
		if *v.Name == name {
			id = *v.Id
			break
		}
	}
	return id, err
}

func (c *Client) GetDepartmentByName(name string) (*Department, error) {
	id, err := c.GetDepartmentIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetDepartment(id)
}

func (c *Client) GetDepartment(id string) (*Department, error) {
	var out *Department
	uri := fmt.Sprintf("%s/%s", uriDepartments, id)

	err := c.doRequest("GET", uri, nil, nil, &out)
	return out, err
}

func (c *Client) GetDepartments() (*ResponseDepartments, error) {
	var out *ResponseDepartments
	err := c.doRequest("GET", uriDepartments, nil, nil, &out)
	return out, err
}

func (c *Client) CreateDepartment(name *string) (*Department, error) {
	in := struct {
		Name *string `json:"name"`
	}{
		Name: name,
	}

	var out *Department

	err := c.doRequest("POST", uriDepartments, in, nil, &out)
	return out, err
}

func (c *Client) UpdateDepartment(d *Department) (*Department, error) {
	var out *Department
	uri := fmt.Sprintf("%s/%s", uriDepartments, *d.Id)

	in := struct {
		Name *string `json:"name"`
	}{
		Name: d.Name,
	}

	err := c.doRequest("PUT", uri, in, nil, &out)
	return out, err
}

func (c *Client) DeleteDepartment(name string) error {
	id, err := c.GetDepartmentIdByName(name)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/%s", uriDepartments, id)
	return c.doRequest("DELETE", uri, nil, nil, nil)
}
