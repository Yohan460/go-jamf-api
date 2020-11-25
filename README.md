# Jamf API in Go

## Install

```go
import "github.com/sioncojp/go-jamf-api"
```

## Usage

sample code: [examples](examples)


## Get auth token in curl
```shell
$ curl -u username:password -X POST "https://xxxxx.jamfcloud.com/uapi/auth/tokens"
$ token=xxxxx
$ curl -X GET -H "Content-Type: application/json" -H "Authorization: Bearer $token" "https://xxxxx.jamfcloud.com/uapi/v1/departments"
```