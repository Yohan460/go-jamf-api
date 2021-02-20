package jamf

type UserScope struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}

type JamfUserScope struct {
	Name string `xml:"name"`
}
