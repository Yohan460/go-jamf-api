package main

import (
	"fmt"
	"os"

	"github.com/sioncojp/go-jamf-api"
)

const (
	// Token ... To get the token in curl, please do the following
	// curl -u username:password -X POST "https://xxxx.jamfcloud.com/uapi/auth/tokens"
	// If your password contains consecutive symbols, your 401 may be returned
	Token = "xxxx"

	// Organization ... xxxx.jamfcloud.com
	Organization = "xxxx"
)

func main() {
	before := "fuga"
	after := "testDepartment"
	c := jamf.NewClient(Token, Organization)

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
