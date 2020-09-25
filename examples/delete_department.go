package main

import (
	"fmt"
	"os"

	"github.com/sioncojp/go-jamf-api"
)

const (
	User     = "xxxx"
	Password = "xxxx"

	// Organization ... xxxx.jamfcloud.com
	Organization = "xxxx"
)

func main() {
	c, err := jamf.NewClient(User, Password, Organization)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := c.DeleteDepartment("fuga"); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d, err := c.GetDepartments()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, v := range d.Results {
		fmt.Printf("%s: %s\n", *v.Id, *v.Name)
	}
}
