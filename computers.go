package jamf

type Computer struct {
	ID           int    `json:"id,omitempty" xml:"id"`
	Name         string `json:"name,omitempty" xml:"name"`
	SerialNumber string `json:"serial_number,omitempty" xml:"serial_number"`
}
