package jamf

type NetworkSegmentScope struct {
	ID   int    `xml:"id"`
	UID  string `xml:"uid,omitempty"`
	Name int    `xml:"name"`
}
