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

	d, err := c.GetDepartment("hoge")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", *d.Id, *d.Name)
}
