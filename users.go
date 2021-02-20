package jamf

type UserScope struct {
	Name string `xml:"name"`
}

type JamfUserScope struct {
	Id   int    `xml:"id"`
	Name string `xml:"name"`
}
