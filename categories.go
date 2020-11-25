package jamf

import "fmt"

const uriCategories = "/v1/categories"

type ResponseCategories struct {
	TotalCount *int       `json:"totalCount,omitempty"`
	Results    []Category `json:"results,omitempty"`
}

type Category struct {
	Id       *string `json:"id,omitempty"` // The response type to be returned is a string
	Name     *string `json:"name,omitempty"`
	Priority *int    `json:"priority,omitempty"`
	Href     *string `json:"href,omitempty"`
}

func (c *Client) GetCategoryIdByName(name string) (string, error) {
	var id string
	d, err := c.GetCategories()
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

func (c *Client) GetCategoryByName(name string) (data *Category, err error) {
	id, err := c.GetCategoryIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetCategory(id)
}

func (c *Client) GetCategory(id string) (data *Category, err error) {
	var out *Category
	uri := fmt.Sprintf("%s/%s", uriCategories, id)

	err = c.doRequest("GET", uri, nil, &out)
	data = out
	return
}

func (c *Client) GetCategories() (data *ResponseCategories, err error) {
	var out *ResponseCategories
	err = c.doRequest("GET", uriCategories, nil, &out)
	data = out
	return
}

func (c *Client) CreateCategory(name *string, priority *int) (*Category, error) {
	in := struct {
		Name     *string `json:"name"`
		Priority *int    `json:"priority"`
	}{
		Name:     name,
		Priority: priority,
	}

	var out *Category

	if err := c.doRequest("POST", uriCategories, in, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) UpdateCategory(d *Category) (*Category, error) {
	var out *Category
	uri := fmt.Sprintf("%s/%s", uriCategories, *d.Id)

	in := struct {
		Name     *string `json:"name"`
		Priority *int    `json:"priority"`
	}{
		Name:     d.Name,
		Priority: d.Priority,
	}

	if err := c.doRequest("PUT", uri, in, &out); err != nil {
		return nil, err
	}
	return out, nil
}

func (c *Client) DeleteCategory(name string) error {
	id, err := c.GetCategoryIdByName(name)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/%s", uriCategories, id)
	return c.doRequest("DELETE", uri, nil, nil)
}
