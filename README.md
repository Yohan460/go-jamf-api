# Jamf API in Go

Mostly implemented Golang library for the [Jamf Pro API](https://developer.jamf.com/jamf-pro/docs/jamf-pro-api-overview).

The Classic API support is written manually, UAPI support is written programmatically using the [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) based upon the publicly available swagger docs.

## Classic API Usage

Example Code

```golang
package main

import (
	"fmt"
	"time"
    "log"

	"github.com/yohan460/go-jamf-api"
)

func main() {
	url := "https://example.jamfcloud.com"
	client, err := jamf.NewClient(
        url,
        jamf.WithOAuth("client-id", "client-secret"),
    )
	if err != nil {
		log.Fatal(err)
	}

	// Use the client to make requests to the Jamf API
}
```

## UAPI Usage

See [Usage docs](api/README.md)

## UAPI Regeneration

When there is a Jamf UAPI schema upgrade the [generate-api.sh](scripts/generate-api.sh) file can be run with the new swagger schema as the first parameter to re-generate it

The script contains a couple of patches to ensure it is syntactically correct go code.
