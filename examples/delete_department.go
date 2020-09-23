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
	Token = "xxx"

	// Organization ... xxxx.jamfcloud.com
	Organization = "xxx"
)

func main() {
	c := jamf.NewClient(Token, Organization)

	err := c.DeleteDepartment("fuga")
	if err != nil {
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
