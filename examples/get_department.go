// In curl:
// curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $token" "$jssurl/uapi/v1/departments"
package main

import (
	"fmt"
	"os"

	"github.com/sioncojp/go-jamf-api"
)

const (
	Username = "xxxx"
	Password = "xxxx"

	// URL ... xxxx.jamfcloud.com
	URL = "xxxx"
)

func main() {
	c, err := jamf.NewClient(Username, Password, URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	d, err := c.GetDepartmentByName("hoge")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%s: %s\n", *d.Id, *d.Name)
}
