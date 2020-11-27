// In curl:
// curl -X DELETE -H "Content-Type: application/json" -H "Authorization: Bearer $token" "$jssurl/uapi/v1/departments/1"
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
