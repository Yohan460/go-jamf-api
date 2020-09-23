package jamf

import "fmt"

const uriDepartments = "/v1/departments"

type ResponseDepartment struct {
	TotalCount *int         `json:"totalCount,omitempty"`
	Results    []Department `json:"results,omitempty"`
}

type Department struct {
	Id   *string `json:"id,omitempty"` // The response type to be returned is a string
	Name *string `json:"name,omitempty"`
	Href *string `json:"href,omitempty"`
}

func (c *Client) DepartmentFindIdByName(name string) (string, error) {
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
	return id, nil
}

func (c *Client) GetDepartment(name string) (data *Department, err error) {
	id, err := c.DepartmentFindIdByName(name)
	if err != nil {
		return nil, err
	}

	var out *Department
	uri := fmt.Sprintf("%s/%s", uriDepartments, id)

	err = c.doRequest("GET", uri, nil, &out)
	data = out
	return
}

func (c *Client) GetDepartments() (data *ResponseDepartment, err error) {
	var out *ResponseDepartment
	err = c.doRequest("GET", uriDepartments, nil, &out)
	data = out
	return
}

func (c *Client) CreateDepartment(name *string) (*Department, error) {
	in := struct {
		Name *string `json:"name"`
	}{
		Name: name,
	}

	var out *Department

	if err := c.doRequest("POST", uriDepartments, in, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) UpdateDepartment(d *Department) (*Department, error) {
	var out *Department
	uri := fmt.Sprintf("%s/%s", uriDepartments, *d.Id)

	in := struct {
		Name *string `json:"name"`
	}{
		Name: d.Name,
	}

	if err := c.doRequest("PUT", uri, in, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) DeleteDepartment(name string) error {
	id, err := c.DepartmentFindIdByName(name)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/%s", uriDepartments, id)
	return c.doRequest("DELETE", uri, nil, nil)
}
