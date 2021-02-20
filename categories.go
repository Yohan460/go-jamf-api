package jamf

import "fmt"

const uriCategories = "/uapi/v1/categories"

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

type GeneralCategory struct {
	ID   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
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
	return id, err
}

func (c *Client) GetCategoryByName(name string) (*Category, error) {
	id, err := c.GetCategoryIdByName(name)
	if err != nil {
		return nil, err
	}

	return c.GetCategory(id)
}

func (c *Client) GetCategory(id string) (*Category, error) {
	var out *Category
	uri := fmt.Sprintf("%s/%s", uriCategories, id)

	err := c.doRequest("GET", uri, nil, &out)
	return out, err
}

func (c *Client) GetCategories() (*ResponseCategories, error) {
	var out *ResponseCategories
	err := c.doRequest("GET", uriCategories, nil, &out)
	return out, err
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

	err := c.doRequest("POST", uriCategories, in, &out)
	return out, err
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

	err := c.doRequest("PUT", uri, in, &out)
	return out, err
}

func (c *Client) DeleteCategory(name string) error {
	id, err := c.GetCategoryIdByName(name)
	if err != nil {
		return err
	}

	uri := fmt.Sprintf("%s/%s", uriCategories, id)
	return c.doRequest("DELETE", uri, nil, nil)
}
