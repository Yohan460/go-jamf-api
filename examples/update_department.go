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
	before := "fuga"
	after := "testDepartment"
	c, err := jamf.NewClient(User, Password, Organization)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := c.GetDepartment(before)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data.Name = &after

	d, err := c.UpdateDepartment(data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", *d.Id, *d.Name)
}
