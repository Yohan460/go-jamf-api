// In curl:
// curl -X PUT -H "Content-Type: application/json" -H "Authorization: Bearer $token" "$jssurl/uapi/v1/departments/1" -d "{\"name\": \"test\"}"
package main

import (
	"fmt"
	"os"

	"github.com/yohan460/go-jamf-api"
)

const (
	Username = "xxxx"
	Password = "xxxx"

	// URL ... xxxx.jamfcloud.com
	URL = "xxxx"
)

func main() {
	before := "fuga"
	after := "testDepartment"
	c, err := jamf.NewClient(Username, Password, URL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	data, err := c.GetDepartmentByName(before)
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
